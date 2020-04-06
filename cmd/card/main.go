package main

import (
	"flag"
	"fmt"
	"github.com/burhon94/CardSvc/cmd/card/server"
	"github.com/burhon94/alifMux/pkg/mux"
	"github.com/burhon94/bdi/pkg/di"
	"log"
	"net"
	"net/http"
)

var (
	host = flag.String("host", "0.0.0.0", "Card service host")
	port = flag.String("port", "10000", "Card service port")
)

func main() {
	addr := net.JoinHostPort(*host, *port)
	serverStart(addr)
}

func serverStart(addr string) {
	container := di.NewContainer()

	if err := container.Provide(
		server.NewServer,
		mux.NewExactMux,
	); err != nil {
		panic(fmt.Sprintf("can't set provide: %v", err))
	}

	container.Start()
	var appServer *server.Server
	container.Component(&appServer)

	log.Printf("Card Service starting on: %s ...", addr)
	panic(http.ListenAndServe(addr, appServer))
}
