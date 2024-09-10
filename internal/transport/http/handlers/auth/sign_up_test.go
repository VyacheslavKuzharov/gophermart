package auth

import (
	"errors"
	"fmt"
	"github.com/VyacheslavKuzharov/gophermart/internal/repository"
	"github.com/VyacheslavKuzharov/gophermart/internal/usecase/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_SignUp(t *testing.T) {
	token := "secretToken"
	login := "qwerty"
	password := "pas12345"

	testCases := map[string]struct {
		useCase        auth.UseCaseInterface
		uri            string
		body           string
		expectedBody   any
		expectedCode   int
		expectedHeader string
		expectedCookie string
	}{
		"200": {
			useCase: &authUseCaseMock{
				SignUpResult: token,
				SignUpError:  nil,
			},
			uri:            "/api/user/register",
			body:           fmt.Sprintf(`{"login": "%s", "password": "%s"}`, login, password),
			expectedBody:   `{}`,
			expectedCode:   http.StatusOK,
			expectedHeader: token,
			expectedCookie: token,
		},
		"400 empty body": {
			useCase: &authUseCaseMock{
				SignUpResult: "",
				SignUpError:  io.EOF,
			},
			uri:            "/api/user/register",
			body:           "",
			expectedBody:   `{"error":"request is empty"}`,
			expectedCode:   http.StatusBadRequest,
			expectedHeader: "",
			expectedCookie: "",
		},
		"400 invalid params": {
			useCase: &authUseCaseMock{
				SignUpResult: "",
				SignUpError:  nil,
			},
			uri:            "/api/user/register",
			body:           fmt.Sprintf(`{"login": "%s", "password": "%s"}`, "", password),
			expectedBody:   `{"error":"login cannot be blank"}`,
			expectedCode:   http.StatusBadRequest,
			expectedHeader: "",
			expectedCookie: "",
		},
		"400 password error": {
			useCase: &authUseCaseMock{
				SignUpResult: "",
				SignUpError:  auth.NewErrGenPwdHash(password, errors.New("test")),
			},
			uri:            "/api/user/register",
			body:           fmt.Sprintf(`{"login": "%s", "password": "%s"}`, login, password),
			expectedBody:   fmt.Sprintf(`{"error":"could not generate password hash from value: %s. test"}`, password),
			expectedCode:   http.StatusBadRequest,
			expectedHeader: "",
			expectedCookie: "",
		},
		"409": {
			useCase: &authUseCaseMock{
				SignUpResult: "",
				SignUpError:  repository.NewUniqueFieldErr(login, errors.New("test")),
			},
			uri:            "/api/user/register",
			body:           fmt.Sprintf(`{"login": "%s", "password": "%s"}`, login, password),
			expectedBody:   fmt.Sprintf(`{"error":"value: %s, already exist."}`, login),
			expectedCode:   http.StatusConflict,
			expectedHeader: "",
			expectedCookie: "",
		},
	}

	for CaseName, test := range testCases {
		t.Run(CaseName, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, test.uri, strings.NewReader(test.body))
			w := httptest.NewRecorder()

			newHandler := New(test.useCase)
			h := newHandler.SignUp
			h(w, r)

			res := w.Result()

			// check response code
			assert.Equal(t, test.expectedCode, w.Code, "response code does not match the expected one")

			// check response body
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			assert.Equal(t, test.expectedBody, strings.TrimSuffix(string(resBody), "\n"))

			// check response Header
			assert.Equal(t, test.expectedHeader, res.Header.Get("Authorization"))

			// check Cookies
			cookies := res.Cookies()
			for _, c := range cookies {
				assert.Equal(t, token, c.Value)
			}
		})
	}
}
