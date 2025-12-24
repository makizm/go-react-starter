package main

import (
	"log/slog"
	"os"

	"go.racktop.io/monorepo/server/internal/server"
)

func main() {
	// @title Monorepo API
	// @version 0.1.0
	// @description API server for the monorepo.
	// @BasePath /api/docs
	srv := server.New(":8080")
	slog.Info("Starting API server", "addr", ":8080")
	if err := srv.ListenAndServe(); err != nil {
		slog.Error("Server failed", "error", err)
		os.Exit(1)
	}
}
