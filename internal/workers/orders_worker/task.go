package ordersworker

import (
	"github.com/VyacheslavKuzharov/gophermart/internal/entity"
	"time"
)

type Task struct {
	entity.Order
	WorkDuration time.Duration
}

func orderToTask(order entity.Order) *Task {
	return &Task{
		Order:        order,
		WorkDuration: 5 * time.Second,
	}
}
