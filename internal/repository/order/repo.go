package order

import (
	"context"
	"github.com/VyacheslavKuzharov/gophermart/internal/entity"
	"github.com/VyacheslavKuzharov/gophermart/pkg/logger"
	"github.com/VyacheslavKuzharov/gophermart/pkg/postgres"
)

type RepositoryInterface interface {
	Create(ctx context.Context, orderDTO *entity.OrderDTO) error
	GetByNumber(ctx context.Context, number string) (entity.Order, error)
	GetManyByStatuses(ctx context.Context, statuses []string) ([]entity.Order, error)
	UpdateStatus(ctx context.Context, status, number string) error
	UpdateAccrualStatus(ctx context.Context, accrual int, status, number string) error
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
