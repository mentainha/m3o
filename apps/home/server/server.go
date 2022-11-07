package server

import (
	"net/http"

	"m3o.dev/apps/home/server/api"
	"m3o.dev/apps/home/server/job"
	"m3o.dev/apps/home/server/web"
)

type Server struct {
	mux *http.ServeMux
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check cors
	if api.Cors(w, r) {
		return
	}
	// serve the default
	s.mux.ServeHTTP(w, r)
}

func New() *Server {
	mux := http.NewServeMux()

	// register web handler
	web.Register(mux)
	// register api routes
	api.Register(mux)
	// register job scheduler
	job.Register()

	return &Server{mux}
}
