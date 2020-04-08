package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

func (s *Server) handlerHealth() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		_, err := writer.Write([]byte("service ok!"))
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

func (s *Server) handlerGetCards() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx, _ := context.WithTimeout(request.Context(), time.Hour)
		cards, err := s.cards.HandleGetCards(ctx)
		if err != nil {
			log.Print(err)
			return
		}

		log.Print(cards)
	}
}

func (s *Server) handlerMyCards() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		ctx, _ := context.WithTimeout(request.Context(), time.Hour)
		myCards, err := s.cards.HandleMyCards(ctx, request)
		if err != nil {
			log.Print(err)
			return
		}

		log.Print(myCards)
	}
}