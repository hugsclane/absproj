package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hugsclane/absproj/go/internal/postgres"
	"go.uber.org/zap"
)

type Config struct {
	Port string
}

type Server struct {
	lg  *zap.Logger
	pg  *postgres.Database
	rd  *redis.Database
	srv *http.Server
}

func NewServer(lg *zap.Logger, pg *postgres.Database, rd *redis.Database, cfg *Config) *Server {
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: mux,
	}
	return &Server{
		lg:  lg,
		pg:  pg,
		rd:  rd,
		srv: srv,
	}
}

func (s *Server) Listen() error {
	return s.srv.ListenAndServe()
}

func (s *Server) GetDataSet(r *http.Request, rw http.ResponseWriter) {
	if ok, msg := s.validateHeaders(r); !ok {
		s.lg.Error("invalid headers", zap.String(msg))
		s.badRequest(rw, msg)
	}

	var req GetDataSetRequest
	if err := readJSON(&req, r.Body); err != nil {
		s.lg.Error("failed to read json body", zap.Error(err))
		s.badRequest(rw, "invalid JSON body")
	}

	if err := s.pg.GetDataSet(req.Key); err != nil {

	}
}

func (s *Server) badRequest(rw http.ResponseWriter, msg string) {
	err := httpError(rw, http.StatusBadRequest, msg)
	if err != nil {
		s.lg.Error("failed to write json body in response")
	}
}

func httpError(rw http.ResponseWriter, code int, msg string) error {
	rw.WriteHeader(code)
	return writeJSON(ErrorResponse{
		Code:    code,
		Message: msg,
	}, rw)
}

func writeJSON(v interface{}, w io.Writer) error {
	enc := json.NewEncoder(w)

	if err := enc.Encode(v); err != nil {
		return fmt.Errorf("failed to write JSON %w", err)
	}

	return nil
}

func readJSON(v interface{}, r io.Reader) error {
	dec := json.NewDecoder(r)

	if err := dec.Decode(v); err != nil {
		return fmt.Errorf("failed to read JSON %w", err)
	}

	return nil
}
