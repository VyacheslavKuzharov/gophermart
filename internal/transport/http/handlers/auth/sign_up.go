package auth

import (
	"encoding/json"
	"errors"
	"github.com/VyacheslavKuzharov/gophermart/internal/lib/response"
	"github.com/VyacheslavKuzharov/gophermart/internal/repository"
	"github.com/VyacheslavKuzharov/gophermart/internal/usecase/auth"
	"io"
	"net/http"
)

type signUpRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req signUpRequest
	var err error
	var genPwdErr *auth.ErrGenPwdHash
	var pgUniqueFieldErr *repository.UniqueFieldErr

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

	err = validateParams(&req)
	if err != nil {
		response.Err(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.useCase.SignUp(ctx, req.Login, req.Password)
	if err != nil {
		if errors.As(err, &genPwdErr) {
			response.Err(w, err.Error(), http.StatusBadRequest)
			return
		}
		if errors.As(err, &pgUniqueFieldErr) {
			response.Err(w, err.Error(), http.StatusConflict)
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
func validateParams(params *signUpRequest) error {
	if params.Login == "" {
		return errors.New("login cannot be blank")
	}

	if params.Password == "" {
		return errors.New("password cannot be blank")
	}

	return nil
}
