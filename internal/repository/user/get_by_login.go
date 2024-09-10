package user

import (
	"context"
	"github.com/VyacheslavKuzharov/gophermart/internal/entity"
	"github.com/VyacheslavKuzharov/gophermart/internal/repository"
)

func (repo *Repository) GetByLogin(ctx context.Context, login string) (entity.User, error) {
	log := repo.logger.Logger
	var user entity.User

	log.Info().Msgf("RUN SQL: SELECT * FROM users WHERE login = %s", login)
	row := repo.pg.Pool.QueryRow(
		ctx,
		"SELECT * FROM users WHERE login = $1",
		login,
	)
	err := row.Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		return entity.User{}, repository.NewNotFountErr("user", "login", login, err)
	}

	return user, nil
}
