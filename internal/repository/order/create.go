package order

import (
	"context"
	"errors"
	"github.com/VyacheslavKuzharov/gophermart/internal/entity"
	"github.com/VyacheslavKuzharov/gophermart/internal/repository"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (repo *Repository) Create(ctx context.Context, orderDTO *entity.OrderDTO) error {
	const target = "repository.order.Create"
	log := repo.logger.Logger
	var orderID int

	log.Info().Msgf("RUN SQL: INSERT INTO orders(user_id, number, status, uploaded_at) VALUES (%s, %s, %s, %s) RETURNING id", orderDTO.UserID, orderDTO.Number, orderDTO.Status, orderDTO.UploadedAt)
	row := repo.pg.Pool.QueryRow(
		ctx,
		"INSERT INTO orders(user_id, number, status, uploaded_at) VALUES ($1, $2, $3, $4) RETURNING id",
		orderDTO.UserID,
		orderDTO.Number,
		orderDTO.Status,
		orderDTO.UploadedAt,
	)
	err := row.Scan(&orderID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			log.Error().Err(err).Msgf("target: %s", target)
			return repository.NewUniqueFieldErr(orderDTO.Number, err)
		}

		log.Error().Err(err).Msgf("target: %s", target)
		return err
	}

	//newUserID, err := uuid.FromString(userUUID)
	//if err != nil {
	//	return nil, err
	//}
	//
	//newUser := &entity.User{
	//	ID:       newUserID,
	//	Login:    userDTO.Login,
	//	Password: userDTO.Password,
	//}
	//
	//return newUser, nil

	return nil
}
