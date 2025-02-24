package accrual

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/VyacheslavKuzharov/gophermart/pkg/logger"
	"net/http"
	"time"
)

type Accrual struct {
	BaseURL string
	Timeout time.Duration
	logger  *logger.Logger
}

func New(addr string, l *logger.Logger) *Accrual {
	a := &Accrual{
		BaseURL: fmt.Sprintf("http://%s/api", addr),
		Timeout: 5 * time.Second,
		logger:  l,
	}

	return a
}

func (a *Accrual) GetOrderInfo(orderNum string) *Response {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Build Request
	url := fmt.Sprintf("%s/orders/%s", a.BaseURL, orderNum)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		a.logger.Logger.Error().Err(err).Msg("NewRequestWithContext")
	}

	// Request Accrual
	var acrRes Response
	client := http.Client{Timeout: a.Timeout}
	resp, err := client.Do(req)
	if err != nil {
		a.logger.Logger.Error().Err(err).Msg("client.Do(req)")
		//return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&acrRes)
	if err != nil {
		a.logger.Logger.Error().Err(err).Msg("decoder.Decode(&acrRes)")
	}

	return &acrRes
}

type Response struct {
	Order   string `json:"order"`
	Status  string `json:"status"`
	Accrual int    `json:"accrual"`
}
