package user

import (
	"context"
	"github.com/VyacheslavKuzharov/gophermart/internal/entity"
	"github.com/VyacheslavKuzharov/gophermart/pkg/logger"
	"github.com/VyacheslavKuzharov/gophermart/pkg/postgres"
)

type RepositoryInterface interface {
	Create(ctx context.Context, user *entity.UserDTO) (*entity.User, error)
	GetByLogin(ctx context.Context, login string) (entity.User, error)
}

type Repository struct {
	pg     *postgres.Pg
	logger *logger.Logger
}

func NewRepo(pg *postgres.Pg, l *logger.Logger) *Repository {
	return &Repository{
		pg:     pg,
		logger: l,
	}
}
