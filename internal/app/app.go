package app

import (
	"fmt"
	"github.com/VyacheslavKuzharov/gophermart/config"
	"github.com/VyacheslavKuzharov/gophermart/internal/di"
	api "github.com/VyacheslavKuzharov/gophermart/internal/transport/http"
	"github.com/VyacheslavKuzharov/gophermart/internal/workers"
	"github.com/VyacheslavKuzharov/gophermart/pkg/httpserver"
	"github.com/VyacheslavKuzharov/gophermart/pkg/logger"
	"github.com/VyacheslavKuzharov/gophermart/pkg/postgres"
	"github.com/go-chi/chi/v5"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	const target = "app.run"

	// Initialize Logger
	l := logger.New(cfg.Log.Level)
	l.Logger.
		Info().
		Str("env", cfg.App.Env).
		Str("version", cfg.App.Version).
		Msg("Initializing Gophermart application...")

	// Initialize Postgres
	pg, err := postgres.New(cfg.PG.DatabaseUri, cfg.PG.PoolMax)
	if err != nil {
		l.Logger.Fatal().Err(err).Msgf("target: %s.postgres.New", target)
	}
	defer pg.Close()

	postgres.RunMigrations(cfg.PG.DatabaseUri, l)

	// Initialize Dependency injection Container
	container := di.NewContainer(pg, l)

	// Initialize Http server
	router := chi.NewRouter()
	api.RegisterRoutes(router, container)

	httpServer := httpserver.New(router, cfg.HTTP.Addr)
	l.Logger.Info().Msgf("Http server started on: %s", cfg.HTTP.Addr)

	// Initialize background workers
	workers.Run(cfg, l)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Logger.Info().Msgf("target: %s signal: %s", target, s.String())
	case err = <-httpServer.Notify():
		l.Logger.Err(fmt.Errorf("target: %s.httpServer.Notify: %w", target, err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Logger.Err(fmt.Errorf("target: %s.httpServer.Shutdown: %w", target, err))
	}
}
