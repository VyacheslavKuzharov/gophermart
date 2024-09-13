package orders

import (
	"fmt"
	"github.com/VyacheslavKuzharov/gophermart/internal/lib/response"
	"io"
	"net/http"
	"strconv"
	"time"
)

type okResponse struct {
	Msg string `json:"msg"`
}

func (h *Handler) Upload(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	//var pgUniqueFieldErr *repository.UniqueFieldErr

	b, err := io.ReadAll(r.Body)
	if err != nil {
		response.Err(w, err.Error(), http.StatusInternalServerError)
		return
	}

	orderNumber := string(b)
	workDuration, _ := time.ParseDuration("5s")
	h.w.QueueTask(orderNumber, workDuration)
	//err = validateOrderNumber(orderNumber)
	//if err != nil {
	//	response.Err(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//
	//err = h.useCase.Upload(ctx, orderNumber)
	//if err != nil {
	//	if errors.As(err, &pgUniqueFieldErr) {
	//		response.OK(w, http.StatusOK, okResponse{Msg: "order already created"})
	//		return
	//	}
	//
	//	response.Err(w, err.Error(), http.StatusBadRequest)
	//	return
	//}

	w.WriteHeader(http.StatusAccepted)
}

func validateOrderNumber(orderNumber string) error {
	//if orderNumber == "" {
	//	return errors.New("order number can't be blank")
	//}

	i, err := strconv.Atoi(orderNumber)
	if err != nil {
		// ... handle error
		panic(err)
	}

	res := Valid(i)
	fmt.Println("----res--->", res)

	return nil
}

func Valid(number int) bool {
	return (number%10+checksum(number/10))%10 == 0
}

func checksum(number int) int {
	var luhn int

	for i := 0; number > 0; i++ {
		cur := number % 10

		if i%2 == 0 { // even
			cur = cur * 2
			if cur > 9 {
				cur = cur%10 + cur/10
			}
		}

		luhn += cur
		number = number / 10
	}
	return luhn % 10
}
