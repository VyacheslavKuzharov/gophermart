package orders

import (
	"context"
	"github.com/VyacheslavKuzharov/gophermart/internal/repository/order"
	"github.com/VyacheslavKuzharov/gophermart/pkg/logger"
)

type UseCaseInterface interface {
	Upload(ctx context.Context, orderNumber string) error
}

type UseCase struct {
	orderRepo order.RepositoryInterface
	logger    *logger.Logger
}

func NewUseCase(orderRepo order.RepositoryInterface, l *logger.Logger) *UseCase {
	return &UseCase{
		orderRepo: orderRepo,
		logger:    l,
	}
}
