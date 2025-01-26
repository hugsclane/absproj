package server

import (
	"net/http"

	"go.uber.org/zap"
)

// GetDataSetRequest - the request for GetDataSet
type GetDataSetRequest struct {
	Key string `json:"key"`
}

// GetDataSetResponse - the response for GetDataSet
type GetDataSetResponse struct {
	Key  string `json:"key"`
	Data string `json:"data"`
}

func (s *Server) GetDataSet(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if ok, code, msg := s.validateHeaders(r); !ok {
		s.lg.Error("invalid headers", zap.String("msg", msg), zap.Int("code", code))
		s.httpError(rw, code, msg)
	}

	var req GetDataSetRequest
	if err := readJSON(r.Body, &req); err != nil {
		s.lg.Error("failed to read json body", zap.Error(err))
		s.badRequest(rw, "invalid JSON body")
	}

	ds, err := s.pg.GetDataSet(ctx, req.Key)
	if err != nil {
		s.lg.Error("failed to get data set", zap.Error(err))
	}

	rs := GetDataSetResponse{
		Key:  ds.Key,
		Data: ds.Data,
	}
	if err := writeJSON(rw, rs); err != nil {
		s.lg.Error("failed to write response data", zap.Error(err))
		s.internalError(rw, "")
	}
}
