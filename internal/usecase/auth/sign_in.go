package auth

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/VyacheslavKuzharov/gophermart/internal/entity"
)

func (uc *UseCase) SignIn(ctx context.Context, login, password string) (*entity.User, error) {
	log := uc.logger.Logger
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(uc.hashSalt))

	user := &entity.User{
		ID:       "1",
		Login:    login,
		Password: fmt.Sprintf("%x", pwd.Sum(nil)),
	}

	log.Info().Msg("SignIn")

	u, err := uc.userRepo.GetByID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return u, nil
}
