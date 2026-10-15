package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/doc-mind-soft/doc-mind/backend/internal/config"
	"github.com/doc-mind-soft/doc-mind/backend/internal/httpapi"
)

const (
	readHeaderTimeout = 5 * time.Second
	readTimeout       = 30 * time.Second
	writeTimeout      = 30 * time.Second
	idleTimeout       = 60 * time.Second
)

func main() {
	cfg := config.Load()

	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	logger := slog.New(logHandler)
	slog.SetDefault(logger)

	httpHandler := httpapi.NewHandler()

	server := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           httpHandler,
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
