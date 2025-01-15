package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/JoshuaWilkes/react-native-example-app/pkg/api"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {

	api := api.API{}

	host := "localhost:8080"

	slog.Info("Starting server", "host", host)
	err := http.ListenAndServe(
		host,
		// For gRPC clients, it's convenient to support HTTP/2 without TLS. You can
		// avoid x/net/http2 by using http.ListenAndServeTLS.
		h2c.NewHandler(api.Handler(), &http2.Server{}),
	)
	slog.Info("Server shutting down")

	return err
}
