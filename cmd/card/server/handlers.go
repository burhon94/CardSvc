package server

import (
	"context"
	"github.com/burhon94/alifMux/pkg/mux"
	"log"
	"net/http"
	"strconv"
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

func (s *Server) handlerMyCard() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx, _ := context.WithTimeout(request.Context(), time.Hour)
		value, ok := mux.FromContext(ctx, "id")
		if !ok {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Printf("can't find from ctx: %v", ok)
			return
		}

		id, err := strconv.Atoi(value)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Printf("can't convert id: %v", err)
			return
		}
		myCard, err := s.cards.HandleMyCard(ctx, request, int64(id))
		if err != nil {
			log.Print(err)
			return
		}

		if myCard.Pan <= 0 {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		log.Print(myCard)
	}
}