//go:generate go tool oapi-codegen -config ./.oapi-codegen.config.yaml ./specification.yaml
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/deifyed/gatekeeper/internal/core"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)

		os.Exit(1)
	}
}

func run() error {
	server := http.Server{
		Addr:    "0.0.0.0:3000",
		Handler: &core.Server{},
	}

	return server.ListenAndServe()
}
