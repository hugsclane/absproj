package server

import "net/http"

type HydratedRoute struct {
	Route
	s *Server
}

type Route struct {
	Get    http.HandlerFunc
	Post   http.HandlerFunc
	Patch  http.HandlerFunc
	Delete http.HandlerFunc
}

func (r *HydratedRoute) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		if r.Get != nil {
			r.Get(rw, req)
		}
	case http.MethodPost:
		if r.Post != nil {
			r.Post(rw, req)
		}
	case http.MethodPatch:
		if r.Patch != nil {
			r.Patch(rw, req)
		}
	case http.MethodDelete:
		if r.Delete != nil {
			r.Delete(rw, req)
		}
	}
	r.s.methodNotAllowed(rw, "method not allowed")
}
