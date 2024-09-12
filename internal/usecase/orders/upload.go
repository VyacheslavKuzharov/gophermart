package orders

import (
	"context"
	"github.com/VyacheslavKuzharov/gophermart/internal/entity"
	uuid "github.com/satori/go.uuid"
	"time"
)

func (uc *UseCase) Upload(ctx context.Context, orderNumber string) error {
	const target = "usecase.orders.Upload"
	currentUserID := ctx.Value(entity.CurrentUserID).(uuid.UUID)
	log := uc.logger.Logger

	orderDTO := &entity.OrderDTO{
		UserID:     currentUserID,
		Number:     orderNumber,
		Status:     "NEW",
		UploadedAt: time.Now(),
	}

	err := uc.orderRepo.Create(ctx, orderDTO)
	if err != nil {
		log.Error().Err(err).Msgf("target: %s.orderRepo.Create", target)
		return err
	}

	return nil
}
