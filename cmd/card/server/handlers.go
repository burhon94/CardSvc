package server

import (
	"context"
	"github.com/burhon94/CardSvc/pkg/core/card"
	"github.com/burhon94/alifMux/pkg/mux"
	readJSON "github.com/burhon94/json/cmd/reader"
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
		ctx, _ := context.WithTimeout(request.Context(), time.Second)
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
		ctx, _ := context.WithTimeout(request.Context(), time.Second)
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
		ctx, _ := context.WithTimeout(request.Context(), time.Second)
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

func (s *Server) handlerCreateCard() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var requestData card.CreateCard
		err := readJSON.ReadJSONHTTP(request, &requestData)
		if err != nil {
			log.Print(err)
			return
		}

		ctx, _ := context.WithTimeout(request.Context(), time.Second)
		err = s.cards.HandleCreateCard(ctx, request, requestData)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Printf("can't create card to client: %v", err)
			return
		}
	}
}

func (s *Server) handlerCardLock() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var requestData card.LockUnLock
		err := readJSON.ReadJSONHTTP(request, &requestData)
		if err != nil {
			log.Print(err)
			return
		}

		ctx, _ := context.WithTimeout(request.Context(), time.Second)
		err = s.cards.HandleCardLock(ctx, request, requestData)
		if err != nil {
			log.Print(err)
			return
		}
	}
}

func (s *Server) handlerCardUnlock() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var requestData card.LockUnLock
		err := readJSON.ReadJSONHTTP(request, &requestData)
		if err != nil {
			log.Print(err)
			return
		}

		ctx, _ := context.WithTimeout(request.Context(), time.Second)
		err = s.cards.HandleCardUnlock(ctx, request, requestData)
		if err != nil {
			log.Print(err)
			return
		}
	}
}
