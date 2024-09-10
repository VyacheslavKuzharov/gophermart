package auth

import (
	"context"
	"github.com/VyacheslavKuzharov/gophermart/internal/entity"
	"github.com/VyacheslavKuzharov/gophermart/internal/lib/token"
	"golang.org/x/crypto/bcrypt"
)

func (uc *UseCase) SignIn(ctx context.Context, login, password string) (string, error) {
	const target = "usecase.auth.SignIn"
	log := uc.logger.Logger

	userDTO := &entity.UserDTO{
		Login:    login,
		Password: password,
	}

	user, err := uc.userRepo.GetByLogin(ctx, userDTO.Login)
	if err != nil {
		log.Error().Err(err).Msgf("target: %s.userRepo.GetByLogin", target)
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDTO.Password))
	if err != nil {
		log.Error().Err(err).Msgf("target: %s.CompareHashAndPassword", target)
		return "", err
	}

	jwt, err := token.CreateJWT(user.ID)
	if err != nil {
		log.Error().Err(err).Msgf("target: %s.createToken", target)
		return "", err
	}

	return jwt, nil
}
