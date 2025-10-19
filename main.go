//go:generate go tool oapi-codegen -config ./.oapi-codegen.config.yaml ./specification.yaml
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/deifyed/gatekeeper/gen"
	"github.com/deifyed/gatekeeper/internal/core"
)

const (
	defaultHost = "0.0.0.0"
	defaultPort = 3000
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)

		os.Exit(1)
	}
}

func run() error {
	host := defaultHost
	port := defaultPort

	handler := gen.HandlerFromMux(&core.Server{}, http.NewServeMux())

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: handler,
	}

	fmt.Printf("Listening on %s:%d\n", host, port)

	return server.ListenAndServe()
}
