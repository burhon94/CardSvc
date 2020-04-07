package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/burhon94/CardSvc/cmd/card/server"
	"github.com/burhon94/CardSvc/pkg/core/card"
	"github.com/burhon94/alifMux/pkg/mux"
	"github.com/burhon94/bdi/pkg/di"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net"
	"net/http"
)

var (
	host = flag.String("host", "0.0.0.0", "Card service host")
	port = flag.String("port", "10000", "Card service port")
	dsn  = flag.String("dsn", "postgres://user:pass@localhost:5566/cards", "Server DSN")
)

type DSN string

func main() {
	addr := net.JoinHostPort(*host, *port)
	serverStart(addr, *dsn)
}

func serverStart(addr string, dsn string) {
	container := di.NewContainer()

	if err := container.Provide(
		server.NewServer,
		mux.NewExactMux,
		card.NewCard,
		func() DSN { return DSN(dsn) },
		func(dsn DSN) *pgxpool.Pool {
			pool, err := pgxpool.Connect(context.Background(), string(dsn))
			if err != nil {
				panic(fmt.Errorf("can't create pool: %w", err))
			}

			return pool
		},
	); err != nil {
		panic(fmt.Sprintf("can't set provide: %v", err))
	}

	container.Start()
	var appServer *server.Server
	container.Component(&appServer)

	log.Printf("Card Service starting on: %s ...", addr)
	panic(http.ListenAndServe(addr, appServer))
}
