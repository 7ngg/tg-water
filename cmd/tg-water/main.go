package main

import (
	"log/slog"
	"os"

	"github.com/7ngg/tg-water/internal/config"
	"github.com/7ngg/tg-water/internal/lib/logger"
	"github.com/7ngg/tg-water/internal/storage/sqlite"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting server", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	_, err := sqlite.NewConnection(cfg.Storage.Path)
	if err != nil {
		log.Error("failed to init storage", logger.Err(err))
		os.Exit(1)
	}

	// TODO: init router: chi, "chi router"

	// TODO: run server
}

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}

	return log
}
