package app

import (
	"fmt"
	"github.com/VyacheslavKuzharov/gophermart/config"
	"github.com/VyacheslavKuzharov/gophermart/pkg/logger"
	"github.com/VyacheslavKuzharov/gophermart/pkg/postgres"
	"html"
	"net/http"
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
		l.Logger.Fatal().Err(err).Msg("app.Run postgres.New")
	}
	defer pg.Close()

	postgres.RunMigrations(cfg.PG.DatabaseUri, l)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/howdy", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Howdy!")
	})

	l.Logger.Fatal().Err(http.ListenAndServe(":8080", nil))
}
