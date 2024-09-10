package auth

import (
	"context"
	"github.com/VyacheslavKuzharov/gophermart/internal/repository/user"
	"github.com/VyacheslavKuzharov/gophermart/pkg/logger"
)

type UseCaseInterface interface {
	SignUp(ctx context.Context, login, password string) (string, error)
	SignIn(ctx context.Context, login, password string) (string, error)
}

type UseCase struct {
	userRepo user.RepositoryInterface
	logger   *logger.Logger
}

func NewUseCase(userRepo user.RepositoryInterface, l *logger.Logger) *UseCase {
	return &UseCase{
		userRepo: userRepo,
		logger:   l,
	}
}
