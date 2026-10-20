package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/doc-mind-soft/doc-mind/backend/internal/config"
	"github.com/doc-mind-soft/doc-mind/backend/internal/httpapi"
)

const (
	readHeaderTimeout = 5 * time.Second
	readTimeout       = 30 * time.Second
	writeTimeout      = 30 * time.Second
	idleTimeout       = 60 * time.Second
	shutdownTimeout   = 10 * time.Second
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

	signalContext, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	serverErrors := make(chan error, 1)

	logger.Info(
		"HTTP server starting",
		slog.String("http_addr", server.Addr),
	)

	go func() {
		serverErrors <- server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error(
				"HTTP server stopped unexpectedly",
				slog.Any("error", err),
			)
			os.Exit(1)
		}

		return

	case <-signalContext.Done():
		logger.Info("shutdown signal received")
	}

	stop()

	shutdownContext, cancel := context.WithTimeout(
		context.Background(),
		shutdownTimeout,
	)
	defer cancel()

	if err := server.Shutdown(shutdownContext); err != nil {
		logger.Error(
			"HTTP server graceful shutdown failed",
			slog.Any("error", err),
		)
		os.Exit(1)
	}

	if err := <-serverErrors; err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error(
			"HTTP server stopped unexpectedly",
			slog.Any("error", err),
		)
		os.Exit(1)
	}

	logger.Info("HTTP server stopped")
}
