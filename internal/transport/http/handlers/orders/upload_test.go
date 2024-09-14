package orders

import (
	"errors"
	"fmt"
	"github.com/VyacheslavKuzharov/gophermart/internal/repository"
	"github.com/VyacheslavKuzharov/gophermart/internal/usecase/orders"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_Upload(t *testing.T) {
	uri := "/api/user/orders"
	validOrderNum := "2730168464161841"
	invalidOrderNum := "2730168464161842"

	var pgUniqueFieldErr *repository.UniqueFieldErr

	testCases := map[string]struct {
		useCase      orders.UseCaseInterface
		uri          string
		body         string
		expectedBody any
		expectedCode int
	}{
		"202": {
			useCase: &ordersUseCaseMock{
				UploadError: nil,
			},
			uri:          uri,
			body:         validOrderNum,
			expectedBody: ``,
			expectedCode: http.StatusAccepted,
		},
		"200": {
			useCase: &ordersUseCaseMock{
				UploadError: pgUniqueFieldErr,
			},
			uri:          uri,
			body:         validOrderNum,
			expectedBody: `{"msg":"order already created"}`,
			expectedCode: http.StatusOK,
		},
		"422": {
			useCase: &ordersUseCaseMock{
				UploadError: nil,
			},
			uri:          uri,
			body:         invalidOrderNum,
			expectedBody: fmt.Sprintf(`{"error":"order number fromat: %s, is invalid"}`, invalidOrderNum),
			expectedCode: http.StatusUnprocessableEntity,
		},
		"400": {
			useCase: &ordersUseCaseMock{
				UploadError: errors.New("test"),
			},
			uri:          uri,
			body:         validOrderNum,
			expectedBody: `{"error":"test"}`,
			expectedCode: http.StatusBadRequest,
		},
		"409": {
			useCase: &ordersUseCaseMock{
				UploadError: repository.NewConflictErr(validOrderNum, errors.New("test")),
			},
			uri:          uri,
			body:         validOrderNum,
			expectedBody: fmt.Sprintf(`{"error":"value: %s, already uploaded by another user"}`, validOrderNum),
			expectedCode: http.StatusConflict,
		},
	}

	for CaseName, test := range testCases {
		t.Run(CaseName, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, test.uri, strings.NewReader(test.body))
			w := httptest.NewRecorder()

			newHandler := New(test.useCase)
			h := newHandler.Upload
			h(w, r)

			res := w.Result()

			// check response code
			assert.Equal(t, test.expectedCode, w.Code, "response code does not match the expected one")

			// check response body
			resBody, err := io.ReadAll(res.Body)
			defer res.Body.Close()

			require.NoError(t, err)
			assert.Equal(t, test.expectedBody, strings.TrimSuffix(string(resBody), "\n"), "response body does not match the expected one")
		})
	}
}
