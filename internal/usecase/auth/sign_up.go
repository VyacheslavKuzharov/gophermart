package auth

import (
	"context"
	"github.com/VyacheslavKuzharov/gophermart/internal/entity"
	"github.com/VyacheslavKuzharov/gophermart/internal/lib/token"
	"golang.org/x/crypto/bcrypt"
)

func (uc *UseCase) SignUp(ctx context.Context, login, password string) (string, error) {
	const target = "usecase.auth.SignUp"
	log := uc.logger.Logger

	pwdBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Error().Err(err).Msgf("target: %s.generateHashPassword", target)
		return "", NewErrGenPwdHash(password, err)
	}

	userDTO := &entity.UserDTO{
		Login:    login,
		Password: string(pwdBytes),
	}

	user, err := uc.userRepo.Create(ctx, userDTO)
	if err != nil {
		log.Error().Err(err).Msgf("target: %s.userRepo.Create", target)
		return "", err
	}

	jwt, err := token.CreateJWT(user.ID)
	if err != nil {
		log.Error().Err(err).Msgf("target: %s.createToken", target)
		return "", err
	}

	return jwt, nil
}
