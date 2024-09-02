package postgres

import (
	"errors"
	"fmt"
	"github.com/VyacheslavKuzharov/gophermart/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"time"
)

const (
	defaultAttempts = 5
	defaultTimeout  = time.Second
)

func RunMigrations(connectURL string, l *logger.Logger) {
	var (
		attempts = defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	l.Logger.Info().Msg("Postgres trying to run migrations...")

	pgURL := fmt.Sprintf(`%s?sslmode=disable`, connectURL)
	for attempts > 0 {
		m, err = migrate.New("file://migrations", pgURL)
		if err == nil {
			break
		}

		l.Logger.Info().Msgf("Migrate: postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(defaultTimeout)
		attempts--
	}
	if err != nil {
		l.Logger.Fatal().Err(err).Msgf("Migrate: postgres connect error: %s", err)
	}

	err = m.Up()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		l.Logger.Fatal().Err(err).Msgf("Migrate: up error: %s", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		l.Logger.Info().Msg("Migrate: no change")
		return
	}

	l.Logger.Info().Msg("Migrate: up success")
}
