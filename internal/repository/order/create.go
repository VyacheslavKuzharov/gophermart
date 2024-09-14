package order

import (
	"context"
	"errors"
	"github.com/VyacheslavKuzharov/gophermart/internal/entity"
	"github.com/VyacheslavKuzharov/gophermart/internal/repository"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	uuid "github.com/satori/go.uuid"
)

func (repo *Repository) Create(ctx context.Context, orderDTO *entity.OrderDTO) error {
	const target = "repository.order.Create"
	log := repo.logger.Logger

	log.Info().Msgf("RUN SQL: INSERT INTO orders(user_id, number, status, uploaded_at) VALUES (%s, %s, %s, %s)", orderDTO.UserID, orderDTO.Number, orderDTO.Status, orderDTO.UploadedAt)
	_, err := repo.pg.Pool.Exec(
		ctx,
		"INSERT INTO orders(user_id, number, status, uploaded_at) VALUES ($1, $2, $3, $4)",
		orderDTO.UserID,
		orderDTO.Number,
		orderDTO.Status,
		orderDTO.UploadedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			log.Error().Err(err).Msgf("target: %s", target)

			existingOrder, _ := repo.GetByNumber(ctx, orderDTO.Number)
			if !uuid.Equal(orderDTO.UserID, existingOrder.UserID) {
				return repository.NewConflictErr(orderDTO.Number, err)
			}

			return repository.NewUniqueFieldErr(orderDTO.Number, err)
		}

		log.Error().Err(err).Msgf("target: %s", target)
		return err
	}

	return nil
}
