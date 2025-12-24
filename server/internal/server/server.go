package server

import (
	"net/http"

	"go.racktop.io/monorepo/server/internal/handlers"
)

// New creates and configures a new HTTP server
func New(addr string) *http.Server {
	mux := http.NewServeMux()
	handlers.RegisterRoutes(mux)
	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}
