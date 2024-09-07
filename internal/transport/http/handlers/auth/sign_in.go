package auth

import (
	"fmt"
	"github.com/go-chi/render"
	"net/http"
)

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------SignIn-->")
	ctx := r.Context()
	//article, ok := ctx.Value("article").(*entity.User)
	//if !ok {
	//	http.Error(w, http.StatusText(422), 422)
	//	return
	//}
	getLogin := "someLogin"
	getPassword := "somePassword"

	user, err := h.useCase.SignIn(ctx, getLogin, getPassword)
	if err != nil {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	render.JSON(w, r, user)
}
