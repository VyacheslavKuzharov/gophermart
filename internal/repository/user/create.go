package user

import (
	"context"
	"errors"
	"github.com/VyacheslavKuzharov/gophermart/internal/entity"
	"github.com/VyacheslavKuzharov/gophermart/internal/repository"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	uuid "github.com/satori/go.uuid"
)

func (repo *Repository) Create(ctx context.Context, userDTO *entity.UserDTO) (*entity.User, error) {
	const target = "repository.user.Create"
	log := repo.logger.Logger
	var userUUID string

	log.Info().Msgf("RUN SQL: INSERT INTO users(login, password) VALUES (%s, %s) RETURNING id", userDTO.Login, userDTO.Password)
	row := repo.pg.Pool.QueryRow(
		ctx,
		"INSERT INTO users(login, password) VALUES ($1, $2) RETURNING id",
		userDTO.Login,
		userDTO.Password,
	)
	err := row.Scan(&userUUID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			log.Error().Err(err).Msgf("target: %s", target)
			return nil, repository.NewUniqueFieldErr(userDTO.Login, err)
		}

		log.Error().Err(err).Msgf("target: %s", target)
		return nil, err
	}

	newUserID, err := uuid.FromString(userUUID)
	if err != nil {
		return nil, err
	}

	newUser := &entity.User{
		ID:       newUserID,
		Login:    userDTO.Login,
		Password: userDTO.Password,
	}

	return newUser, nil
}
