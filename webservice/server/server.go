package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/hugsclane/absproj/webservice/internal/postgres"
	"github.com/hugsclane/absproj/webservice/internal/redis"
	"go.uber.org/zap"
)

type Config struct {
	Port int
}

type Server struct {
	lg  *zap.Logger
	pg  *postgres.Database
	rd  *redis.Database
	srv *http.Server
}

func NewServer(lg *zap.Logger, pg *postgres.Database, rd *redis.Database, cfg Config) *Server {
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: mux,
	}
	s := &Server{
		lg:  lg,
		pg:  pg,
		rd:  rd,
		srv: srv,
	}

	mux.Handle("/dataset", s.hydrate(Route{
		Get: s.GetDataSet,
	}))

	return s
}

func (s *Server) Listen() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Close() error {
	// drain connections for up to 5s and close
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return s.srv.Shutdown(ctx)
}

func (s *Server) hydrate(r Route) *HydratedRoute {
	return &HydratedRoute{Route: r, s: s}
}

func (s *Server) validateHeaders(r *http.Request) (bool, int, string) {
	if ok, code, msg := validateAccept(r); !ok {
		return ok, code, msg
	}
	if ok, code, msg := validateContentType(r); !ok {
		return ok, code, msg
	}

	return true, 0, ""
}

func validateAccept(r *http.Request) (bool, int, string) {
	if !validateApplicationMetaType(r.Header.Get("Accept"), "json") {
		return false, http.StatusNotAcceptable, "invalid accept header, application/json required"
	}

	return true, 0, ""
}

func validateContentType(r *http.Request) (bool, int, string) {
	if !validateApplicationMetaType(r.Header.Get("Content-Type"), "json") {
		return false, http.StatusUnsupportedMediaType, "invalid content-type, application/json required"
	}

	return true, 0, ""
}

func validateApplicationMetaType(src string, tp string) bool {
	metaTypes := strings.Split(src, ",")

	primary := slices.Contains(metaTypes, "application/"+tp)

	wild1 := slices.Contains(metaTypes, "application/*")

	wild2 := slices.Contains(metaTypes, "*/*")

	return primary || wild1 || wild2
}

func (s *Server) methodNotAllowed(rw http.ResponseWriter, msg string) {
	s.httpError(rw, http.StatusMethodNotAllowed, msg)
}

func (s *Server) badRequest(rw http.ResponseWriter, msg string) {
	s.httpError(rw, http.StatusBadRequest, msg)
}

func (s *Server) internalError(rw http.ResponseWriter, msg string) {
	s.httpError(rw, http.StatusInternalServerError, msg)
}

func (s *Server) httpError(rw http.ResponseWriter, code int, msg string) {
	rw.WriteHeader(code)
	err := writeJSON(rw, ErrorResponse{
		Code:    code,
		Message: msg,
	})

	if err != nil {
		s.lg.Error("failed to write json body in response")
	}
}

func writeJSON(w io.Writer, v interface{}) error {
	enc := json.NewEncoder(w)

	if err := enc.Encode(v); err != nil {
		return fmt.Errorf("failed to write JSON %w", err)
	}

	return nil
}

func readJSON(r io.Reader, v interface{}) error {
	dec := json.NewDecoder(r)

	if err := dec.Decode(v); err != nil {
		return fmt.Errorf("failed to read JSON %w", err)
	}

	return nil
}
