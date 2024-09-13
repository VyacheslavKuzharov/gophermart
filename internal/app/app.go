package app

import (
	"context"
	"fmt"
	"github.com/VyacheslavKuzharov/gophermart/config"
	"github.com/VyacheslavKuzharov/gophermart/internal/di"
	"github.com/VyacheslavKuzharov/gophermart/internal/scheduler"
	api "github.com/VyacheslavKuzharov/gophermart/internal/transport/http"
	"github.com/VyacheslavKuzharov/gophermart/internal/workers"
	"github.com/VyacheslavKuzharov/gophermart/pkg/httpserver"
	"github.com/VyacheslavKuzharov/gophermart/pkg/logger"
	"github.com/VyacheslavKuzharov/gophermart/pkg/postgres"
	"github.com/go-chi/chi/v5"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func Run(cfg *config.Config) {
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
		l.Logger.Fatal().Err(err).Msg("app.Run - postgres.New")
	}
	defer pg.Close()

	postgres.RunMigrations(cfg.PG.DatabaseUri, l)

	// Initialize Dependency injection Container
	container := di.NewContainer(pg, l)

	// Initialize Http server
	router := chi.NewRouter()

	w := workers.New(3, 2)
	l.Logger.Info().Msg("Starting workers.....")
	w.Start(context.Background())

	jobScheduler := scheduler.NewJobScheduler(1 * time.Second)
	jobScheduler.Start()

	job := scheduler.PrintJob{Message: "-----------Hello, World!----------->"}
	jobScheduler.JobQueue <- job

	api.RegisterRoutes(router, container, w)

	httpServer := httpserver.New(router, cfg.HTTP.Addr)
	l.Logger.Info().Msgf("Http server started on: %s", cfg.HTTP.Addr)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	fmt.Printf("Number of goroutines main: %d\n", runtime.NumGoroutine())

	select {
	case s := <-interrupt:
		l.Logger.Info().Msgf("app.Run - signal: %s", s.String())
	case err = <-httpServer.Notify():
		l.Logger.Err(fmt.Errorf("app.Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Logger.Err(fmt.Errorf("app.Run - httpServer.Shutdown: %w", err))
	}
}
