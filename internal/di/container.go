package di

import (
	"github.com/VyacheslavKuzharov/gophermart/config"
	"github.com/VyacheslavKuzharov/gophermart/internal/repository/order"
	"github.com/VyacheslavKuzharov/gophermart/internal/repository/user"
	"github.com/VyacheslavKuzharov/gophermart/internal/usecase/auth"
	"github.com/VyacheslavKuzharov/gophermart/internal/usecase/orders"
	ordersworker "github.com/VyacheslavKuzharov/gophermart/internal/workers/orders_worker"
	"github.com/VyacheslavKuzharov/gophermart/pkg/logger"
	"github.com/VyacheslavKuzharov/gophermart/pkg/postgres"
	"time"
)

type Container struct {
	Pg            *postgres.Pg
	Logger        *logger.Logger
	cfg           *config.Config
	orderRepo     *order.Repository
	userRepo      *user.Repository
	ordersUseCase *orders.UseCase
	authUseCase   *auth.UseCase
	ordersWorker  *ordersworker.OrdersWorker
}

func NewContainer(pg *postgres.Pg, cfg *config.Config, l *logger.Logger) *Container {
	return &Container{
		Pg:     pg,
		Logger: l,
		cfg:    cfg,
	}
}

func (c *Container) GetAuthUseCase() *auth.UseCase {
	if c.authUseCase == nil {
		c.authUseCase = auth.NewUseCase(c.getUserRepo(), c.Logger)
	}

	return c.authUseCase
}

func (c *Container) GetOrdersUseCase() *orders.UseCase {
	if c.ordersUseCase == nil {
		c.ordersUseCase = orders.NewUseCase(c.getOrderRepo(), c.Logger)
	}

	return c.ordersUseCase
}

func (c *Container) GetOrdersWorker() *ordersworker.OrdersWorker {
	if c.ordersWorker == nil {
		poling := time.Duration(c.cfg.Worker.Poling) * time.Second

		c.ordersWorker = ordersworker.New(poling, c.getOrderRepo(), c.Logger, c.cfg)
	}

	return c.ordersWorker
}

func (c *Container) getUserRepo() *user.Repository {
	if c.userRepo == nil {
		c.userRepo = user.NewRepo(c.Pg, c.Logger)
	}

	return c.userRepo
}

func (c *Container) getOrderRepo() *order.Repository {
	if c.orderRepo == nil {
		c.orderRepo = order.NewRepo(c.Pg, c.Logger)
	}

	return c.orderRepo
}
