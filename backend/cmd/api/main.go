package main

import (
	"log/slog"
	"os"

	"github.com/doc-mind-soft/doc-mind/backend/internal/config"
)

func main() {
	cfg := config.Load()

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	logger.Info(
		"application initialized",
		slog.String("http_addr", cfg.HTTPAddr),
	)
}
