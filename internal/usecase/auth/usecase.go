package auth

import (
	"context"
	"github.com/VyacheslavKuzharov/gophermart/internal/entity"
	"github.com/VyacheslavKuzharov/gophermart/internal/repository/user"
	"github.com/VyacheslavKuzharov/gophermart/pkg/logger"
)

type UseCaseInterface interface {
	SignUp(ctx context.Context, username, password string) (*entity.User, error)
	SignIn(ctx context.Context, username, password string) (*entity.User, error)
}

type UseCase struct {
	logger   *logger.Logger
	userRepo user.RepositoryInterface
	hashSalt string
}

func NewUseCase(l *logger.Logger, userRepo user.RepositoryInterface, hashSalt string) *UseCase {
	return &UseCase{
		logger:   l,
		userRepo: userRepo,
		hashSalt: hashSalt,
	}
}
