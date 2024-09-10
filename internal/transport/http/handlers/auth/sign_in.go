package auth

import (
	"encoding/json"
	"errors"
	"github.com/VyacheslavKuzharov/gophermart/internal/lib/response"
	"github.com/VyacheslavKuzharov/gophermart/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
)

type signInRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req signInRequest
	var err error
	var notFoundErr *repository.NotFountErr

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if errors.Is(err, io.EOF) {
		response.Err(w, "request is empty", http.StatusBadRequest)
		return
	}
	if err != nil {
		response.Err(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = validateSignInParams(&req)
	if err != nil {
		response.Err(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.useCase.SignIn(ctx, req.Login, req.Password)
	if err != nil {
		if errors.As(err, &notFoundErr) {
			response.Err(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			response.Err(w, "invalid password", http.StatusUnauthorized)
			return
		}

		response.Err(w, err.Error(), http.StatusBadRequest)
		return
	}

	cookie := http.Cookie{
		Name:  "Authorization",
		Value: token,
		Path:  "/",
	}

	http.SetCookie(w, &cookie)
	w.Header().Set(cookie.Name, cookie.Value)

	response.OK(w, http.StatusOK, struct{}{})
}

// TODO: Added more validations
func validateSignInParams(params *signInRequest) error {
	if params.Login == "" {
		return errors.New("login cannot be blank")
	}

	if params.Password == "" {
		return errors.New("password cannot be blank")
	}

	return nil
}
