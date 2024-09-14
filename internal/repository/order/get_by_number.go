package order

import (
	"context"
	"github.com/VyacheslavKuzharov/gophermart/internal/entity"
	"github.com/VyacheslavKuzharov/gophermart/internal/repository"
	uuid "github.com/satori/go.uuid"
)

func (repo *Repository) GetByNumber(ctx context.Context, number string) (entity.Order, error) {
	log := repo.logger.Logger
	var order entity.Order
	var userIDStr string

	log.Info().Msgf("RUN SQL: SELECT * FROM orders WHERE number = %s", number)
	row := repo.pg.Pool.QueryRow(
		ctx,
		`SELECT * FROM orders WHERE number = $1`,
		number,
	)

	err := row.Scan(&order.ID, &userIDStr, &order.Number, &order.Status, &order.Accrual, &order.UploadedAt)
	if err != nil {
		return entity.Order{}, repository.NewNotFountErr("order", "number", number, err)
	}

	userID, _ := uuid.FromString(userIDStr)
	order.UserID = userID

	return order, nil
}
