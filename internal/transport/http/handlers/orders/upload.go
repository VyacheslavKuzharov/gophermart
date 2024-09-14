package orders

import (
	"errors"
	"fmt"
	luhncheck "github.com/VyacheslavKuzharov/gophermart/internal/lib/luhn_check"
	"github.com/VyacheslavKuzharov/gophermart/internal/lib/response"
	"github.com/VyacheslavKuzharov/gophermart/internal/repository"
	"io"
	"net/http"
	"strconv"
)

type uploadResp struct {
	Msg string `json:"msg"`
}

func (h *Handler) Upload(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var pgUniqueFieldErr *repository.UniqueFieldErr
	var pgConflictErr *repository.ConflictErr

	b, err := io.ReadAll(r.Body)
	if err != nil {
		response.Err(w, err.Error(), http.StatusInternalServerError)
		return
	}

	orderNumber := string(b)
	err = validateOrderNumber(orderNumber)
	if err != nil {
		response.Err(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = h.useCase.Upload(ctx, orderNumber)
	if err != nil {
		if errors.As(err, &pgUniqueFieldErr) {
			response.OK(w, http.StatusOK, uploadResp{Msg: "order already created"})
			return
		}
		if errors.As(err, &pgConflictErr) {
			response.Err(w, err.Error(), http.StatusConflict)
			return
		}

		response.Err(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func validateOrderNumber(orderNumber string) error {
	if orderNumber == "" {
		return errors.New("order number can't be blank")
	}

	num, err := strconv.Atoi(orderNumber)
	if err != nil {
		return err
	}

	if !luhncheck.Valid(num) {
		return fmt.Errorf("order number fromat: %d, is invalid", num)
	}

	return nil
}
