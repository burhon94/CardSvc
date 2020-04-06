package main

import (
	"net"
	"net/http"
)

func main() {
	host := "0.0.0.0"
	port := "10000"
	addr := net.JoinHostPort(host, port)

	serverStart(addr)
}

func serverStart(addr string) {
	panic(http.ListenAndServe(addr, indexHandler()))
}

func indexHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		_, err := writer.Write([]byte("Hello from Card Service"))
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
