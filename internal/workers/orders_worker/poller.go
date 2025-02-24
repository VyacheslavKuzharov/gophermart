package ordersworker

import (
	"context"
	"github.com/VyacheslavKuzharov/gophermart/internal/entity"
	"github.com/VyacheslavKuzharov/gophermart/internal/entity/order"
	"time"
)

func (ow *OrdersWorker) spawnDBPoller(ctx context.Context) {
	defer ow.wg.Done()

	ticker := time.NewTicker(ow.Polling)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			orders := ow.fetchAvailableOrders(ctx)

			if len(orders) > 0 {
				for _, order := range orders {
					ow.enqueueTask(orderToTask(order))
				}
			}
		}
	}
}

func (ow *OrdersWorker) fetchAvailableOrders(ctx context.Context) []entity.Order {
	const target = "ordersworker.poller.fetchAvailableOrders"
	log := ow.logger.Logger
	targetStatuses := []string{order.StatusNew.String(), order.StatusProcessing.String()}

	orders, err := ow.repo.GetManyByStatuses(ctx, targetStatuses)
	// TODO: подумачть что здесь делать
	if err != nil {
		log.Error().Err(err).Msgf("target: %s.GetManyByStatuses", target)
	}

	return orders
}
