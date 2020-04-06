package server

import (
	"github.com/burhon94/alifMux/pkg/mux"
	"net/http"
)

type Server struct {
	router *mux.ExactMux
}

func NewServer(router *mux.ExactMux) *Server {
	return &Server{router: router}
}

func (s *Server) Start()  {
	s.InitRoutes()
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	s.router.ServeHTTP(w, r)
}
