package auth

import (
	"context"
	"github.com/VyacheslavKuzharov/gophermart/internal/entity"
	logdiscard "github.com/VyacheslavKuzharov/gophermart/internal/lib/log_discard"
	"github.com/VyacheslavKuzharov/gophermart/internal/lib/token"
	"github.com/VyacheslavKuzharov/gophermart/internal/repository"
	"github.com/VyacheslavKuzharov/gophermart/internal/repository/user"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestUseCase_SignIn(t *testing.T) {
	log := logdiscard.NewDiscardLogger()
	login := "qwerty"
	password := "password"

	var notFoundErr *repository.NotFountErr

	pwdBytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	existingUser := entity.User{
		ID:       uuid.NewV4(),
		Login:    login,
		Password: string(pwdBytes),
	}

	testCases := map[string]struct {
		repo           user.RepositoryInterface
		expectedResult string
		expectedErr    error
		paramsLogin    string
		paramsPwd      string
	}{
		"Happy path": {
			repo: &userRepoMock{
				GetByLoginResult: existingUser,
				GetByLoginError:  nil,
			},
			paramsLogin:    login,
			paramsPwd:      password,
			expectedResult: existingUser.ID.String(),
			expectedErr:    nil,
		},
		"Login not found": {
			repo: &userRepoMock{
				GetByLoginResult: entity.User{},
				GetByLoginError:  notFoundErr,
			},
			paramsLogin:    "invalidLogin",
			paramsPwd:      password,
			expectedResult: "",
			expectedErr:    notFoundErr,
		},
		"Password not match": {
			repo: &userRepoMock{
				GetByLoginResult: existingUser,
				GetByLoginError:  nil,
			},
			paramsLogin:    login,
			paramsPwd:      "invalid",
			expectedResult: "",
			expectedErr:    bcrypt.ErrMismatchedHashAndPassword,
		},
	}

	for CaseName, test := range testCases {
		t.Run(CaseName, func(t *testing.T) {

			usecase := NewUseCase(test.repo, log)
			jwt, err := usecase.SignIn(context.Background(), test.paramsLogin, test.paramsPwd)

			if err == nil {
				claims, _ := token.ParseJWT(jwt)
				assert.Equal(t, test.expectedResult, claims.UserID.String())
			}

			assert.Equal(t, test.expectedErr, err)
		})
	}
}
