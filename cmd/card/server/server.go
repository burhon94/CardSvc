package server

import (
	"github.com/burhon94/CardSvc/pkg/core/card"
	"github.com/burhon94/alifMux/pkg/mux"
	"net/http"
)

type Server struct {
	router *mux.ExactMux
	cards  *card.Card
}

func NewServer(router *mux.ExactMux, cards *card.Card) *Server {
	return &Server{router: router, cards: cards}
}

func (s *Server) Start() {
	s.InitRoutes()
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
