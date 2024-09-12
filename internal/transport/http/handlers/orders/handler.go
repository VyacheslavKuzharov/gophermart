package orders

import (
	"github.com/VyacheslavKuzharov/gophermart/internal/usecase/orders"
)

type Handler struct {
	useCase orders.UseCaseInterface
}

func New(useCase orders.UseCaseInterface) *Handler {
	return &Handler{
		useCase: useCase,
	}
}
