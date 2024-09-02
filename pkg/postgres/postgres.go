package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Pg struct {
	maxPoolSize int
	Pool        *pgxpool.Pool
}

func New(connectURL string, maxPoolSize int) (*Pg, error) {
	pg := &Pg{
		maxPoolSize: maxPoolSize,
	}

	poolConfig, err := pgxpool.ParseConfig(connectURL)
	if err != nil {
		return nil, fmt.Errorf("postgres. pgxpool.ParseConfig: %w", err)
	}

	poolConfig.MaxConns = int32(pg.maxPoolSize)

	pg.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("postgres. NewPostgres connError: %w", err)
	}

	return pg, nil
}

// Close - Closes Pg connection
func (p *Pg) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
