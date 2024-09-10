package auth

import (
	"context"
	"github.com/VyacheslavKuzharov/gophermart/internal/entity"
)

type userRepoMock struct {
	CreateResult *entity.User
	CreateError  error

	GetByLoginResult entity.User
	GetByLoginError  error
}

func (repo *userRepoMock) Create(ctx context.Context, user *entity.UserDTO) (*entity.User, error) {
	return repo.CreateResult, repo.CreateError
}

func (repo *userRepoMock) GetByLogin(ctx context.Context, login string) (entity.User, error) {
	return repo.GetByLoginResult, repo.GetByLoginError
}
