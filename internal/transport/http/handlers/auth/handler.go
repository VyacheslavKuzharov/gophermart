package auth

import "github.com/VyacheslavKuzharov/gophermart/internal/usecase/auth"

type Handler struct {
	useCase auth.UseCaseInterface
}

func New(useCase auth.UseCaseInterface) *Handler {
	return &Handler{
		useCase: useCase,
	}
}
