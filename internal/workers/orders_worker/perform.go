package ordersworker

import (
	"context"
	"fmt"
	"github.com/VyacheslavKuzharov/gophermart/internal/entity/order"
)

func (ow *OrdersWorker) perform(ctx context.Context, task *Task) {
	const target = "ordersworker.perform"
	log := ow.logger.Logger

	log.Info().Msgf("Worker queue task with order number: %s", task.Order.Number)

	if task.Order.Status == order.StatusNew.String() {
		err := ow.repo.UpdateStatus(ctx, order.StatusProcessing.String(), task.Order.Number)
		if err != nil {
			log.Error().Err(err).Msgf("target: %s.UpdateStatus", target)
		}
	}

	orderNumInfo := ow.accrual.GetOrderInfo(task.Order.Number)

	fmt.Println("----orderInfo---->", orderNumInfo)

	fmt.Println("worker update order record with accrual")
}
