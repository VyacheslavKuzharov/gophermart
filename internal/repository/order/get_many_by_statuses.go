package order

import (
	"context"
	"github.com/VyacheslavKuzharov/gophermart/internal/entity"
	uuid "github.com/satori/go.uuid"
	"strings"
)

func (repo *Repository) GetManyByStatuses(ctx context.Context, statuses []string) ([]entity.Order, error) {
	log := repo.logger.Logger
	var orders []entity.Order

	statusesStr := `'` + strings.Join(statuses, `', '`) + `'`

	log.Info().Msgf("RUN SQL: SELECT * FROM orders WHERE staus IN (%s)", statusesStr)
	rows, err := repo.pg.Pool.Query(
		ctx,
		`SELECT * FROM orders WHERE status IN(`+statusesStr+`);`,
	)
	if err != nil {
		return []entity.Order{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var userIDStr string
		var ord entity.Order

		err = rows.Scan(&ord.ID, &userIDStr, &ord.Number, &ord.Status, &ord.Accrual, &ord.UploadedAt)
		if err != nil {
			return []entity.Order{}, err
		}

		userID, _ := uuid.FromString(userIDStr)
		ord.UserID = userID

		orders = append(orders, ord)
	}

	return orders, nil
}
