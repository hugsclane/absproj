package server

import (
	"net/http"

	"github.com/hugsclane/absproj/go/internal/model"
	"github.com/hugsclane/absproj/go/internal/redis"
	"go.uber.org/zap"
)

// GetDataSetRequest - the request for GetDataSet
type GetDataSetRequest struct {
	Key string `json:"key"`
}

// GetDataSetResponse - the response for GetDataSet
type GetDataSetResponse struct {
	Key  string `json:"key"`
	Data []byte `json:"data"`
}

func (s *Server) GetDataSet(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sendResp := func(ds *model.DataSet) {
		rs := GetDataSetResponse{
			Key:  ds.Key,
			Data: ds.Data,
		}
		if err := writeJSON(rw, rs); err != nil {
			s.lg.Error("failed to write response data", zap.Error(err))
			s.internalError(rw, "")
		}
	}

	// validate headers
	if ok, code, msg := s.validateHeaders(r); !ok {
		s.lg.Error("invalid headers", zap.String("msg", msg), zap.Int("code", code))
		s.httpError(rw, code, msg)
	}

	// decode body
	var req GetDataSetRequest
	if err := readJSON(r.Body, &req); err != nil {
		s.lg.Error("failed to read json body", zap.Error(err))
		s.badRequest(rw, "invalid JSON body")
	}

	// check cache
	ds, err := s.rd.GetDataSet(ctx, req.Key)
	if err == nil {
		sendResp(ds)
		return
	}
	if err != nil {
		if err == redis.ErrNotExists {
			s.lg.Info("data set cache miss", zap.String("key", ds.Key))
		}
		s.lg.Error("failed to get data set", zap.Error(err))
	}

	// check database
	ds, err = s.pg.GetDataSet(ctx, req.Key)
	if err != nil {
		s.lg.Error("failed to get data set", zap.Error(err))
	}

	// async set the cache
	go func() {
		if err := s.rd.AddDataSet(ctx, ds); err != nil {
			s.lg.Error("failed to add data set to redis async: %w", zap.Error(err))
		}
	}()
}
