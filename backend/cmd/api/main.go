package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/doc-mind-soft/doc-mind/backend/internal/config"
)

const (
	readHeaderTimeout = 5 * time.Second
	readTimeout       = 30 * time.Second
	writeTimeout      = 30 * time.Second
	idleTimeout       = 60 * time.Second
)

func main() {
	cfg := config.Load()

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           mux,
		ReadHeaderTimeout: readHeaderTimeout,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
	}

	logger.Info(
		"HTTP server starting",
		slog.String("http_addr", server.Addr),
	)

	if err := server.ListenAndServe(); err != nil {
		logger.Error(
			"HTTP server stopped",
			slog.Any("error", err),
		)
		os.Exit(1)
	}
}
