package di

import (
	"github.com/VyacheslavKuzharov/gophermart/internal/repository/user"
	"github.com/VyacheslavKuzharov/gophermart/internal/usecase/auth"
	"github.com/VyacheslavKuzharov/gophermart/pkg/logger"
	"github.com/VyacheslavKuzharov/gophermart/pkg/postgres"
)

type Container struct {
	pg     *postgres.Pg
	logger *logger.Logger
}

func NewContainer(pg *postgres.Pg, l *logger.Logger) *Container {
	return &Container{
		pg:     pg,
		logger: l,
	}
}

func (c *Container) GetAuthUseCase() *auth.UseCase {
	return auth.NewUseCase(c.logger, c.getUserRepo(), "qwerty")
}

func (c *Container) getUserRepo() *user.Repository {
	return user.NewRepo(c.pg)
}
