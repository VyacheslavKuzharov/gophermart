package user

import (
	"context"
	"github.com/VyacheslavKuzharov/gophermart/internal/entity"
	"github.com/VyacheslavKuzharov/gophermart/pkg/postgres"
)

type RepositoryInterface interface {
	GetByID(ctx context.Context, id string) (*entity.User, error)
}

type Repository struct {
	pg *postgres.Pg
}

func NewRepo(pg *postgres.Pg) *Repository {
	return &Repository{
		pg: pg,
	}
}
