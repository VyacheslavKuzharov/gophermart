package orders

import (
	"github.com/VyacheslavKuzharov/gophermart/internal/usecase/orders"
	"github.com/VyacheslavKuzharov/gophermart/internal/workers"
)

type Handler struct {
	useCase orders.UseCaseInterface
	w       workers.WorkerIface
}

func New(useCase orders.UseCaseInterface, w workers.WorkerIface) *Handler {
	return &Handler{
		useCase: useCase,
		w:       w,
	}
}
