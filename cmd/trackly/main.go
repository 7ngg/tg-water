package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/7ngg/trackly/internal/config"
	"github.com/7ngg/trackly/internal/http-server/handlers"
	"github.com/7ngg/trackly/internal/lib/logger"
	"github.com/7ngg/trackly/internal/storage/sqlite"
	"github.com/7ngg/trackly/internal/trackly"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	log.Info("starting server", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	storage, err := sqlite.NewConnection(cfg.Storage.Path)
	if err != nil {
		log.Error("failed to init storage", logger.Err(err))
		os.Exit(1)
	}

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(cfg.HTTPServer.Timeout))

	router.Route("/users", func(r chi.Router) {
		r.Get("/", handlers.GetAllUsers(log, storage.DB))
	})

	srv := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("failed to start server", logger.Err(err))
			return
		}
	}()

	log.Debug("server started")

	b, err := trackly.New(cfg.BotToken, storage, log)
	if err != nil {
		log.Error("failed to initialize bot", logger.Err(err))
		return
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b.Bot.Start(ctx)
}
