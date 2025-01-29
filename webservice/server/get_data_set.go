package server

import (
	"fmt"
	"net/http"

	"github.com/hugsclane/absproj/webservice/internal/model"
	"github.com/hugsclane/absproj/webservice/internal/postgres"
	"github.com/hugsclane/absproj/webservice/internal/redis"
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

	intErr := func() {
		s.internalError(rw, "internal server error while processing request")
	}

	sendResp := func(ds *model.DataSet) {
		rs := GetDataSetResponse{
			Key:  ds.Key,
			Data: ds.Data,
		}
		if err := writeJSON(rw, rs); err != nil {
			s.lg.Error("failed to write response data", zap.Error(err))
			intErr()
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

	if err != redis.ErrNotExists {
		s.lg.Error("failed to get data set", zap.Error(err))
		intErr()
		return
	}

	s.lg.Info("data set cache miss", zap.String("key", ds.Key))

	// check database
	ds, err = s.pg.GetDataSet(ctx, req.Key)
	if err == postgres.ErrNotExists {
		s.httpError(rw, http.StatusNotFound, fmt.Sprintf("the data set does not exist"))
		return
	}
	if err != nil {
		s.lg.Error("failed to get data set", zap.Error(err))
		intErr()
		return
	}

	// async set the cache
	go func() {
		if err := s.rd.AddDataSet(ctx, ds); err != nil {
			s.lg.Error("failed to add data set to redis async: %w", zap.Error(err))
		}
	}()

	sendResp(ds)
}
