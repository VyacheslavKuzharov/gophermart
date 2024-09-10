package di

import (
	"github.com/VyacheslavKuzharov/gophermart/internal/repository/user"
	"github.com/VyacheslavKuzharov/gophermart/internal/usecase/auth"
	"github.com/VyacheslavKuzharov/gophermart/pkg/logger"
	"github.com/VyacheslavKuzharov/gophermart/pkg/postgres"
)

type Container struct {
	Pg     *postgres.Pg
	Logger *logger.Logger
}

func NewContainer(pg *postgres.Pg, l *logger.Logger) *Container {
	return &Container{
		Pg:     pg,
		Logger: l,
	}
}

func (c *Container) GetAuthUseCase() *auth.UseCase {
	return auth.NewUseCase(c.getUserRepo(), c.Logger)
}

func (c *Container) getUserRepo() *user.Repository {
	return user.NewRepo(c.Pg, c.Logger)
}
