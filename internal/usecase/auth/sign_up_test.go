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
	"testing"
)

func TestUseCase_SignUp(t *testing.T) {
	log := logdiscard.NewDiscardLogger()
	var pgUniqueFieldErr *repository.UniqueFieldErr

	newUser := &entity.User{
		ID:       uuid.NewV4(),
		Login:    "qwerty",
		Password: "password",
	}

	testCases := map[string]struct {
		repo           user.RepositoryInterface
		expectedResult string
		expectedErr    error
	}{
		"Happy path": {
			repo: &userRepoMock{
				CreateResult: newUser,
				CreateError:  nil,
			},
			expectedResult: newUser.ID.String(),
			expectedErr:    nil,
		},
		"Login already exists": {
			repo: &userRepoMock{
				CreateResult: nil,
				CreateError:  pgUniqueFieldErr,
			},
			expectedResult: "",
			expectedErr:    pgUniqueFieldErr,
		},
	}

	for CaseName, test := range testCases {
		t.Run(CaseName, func(t *testing.T) {

			usecase := NewUseCase(test.repo, log)
			jwt, err := usecase.SignUp(context.Background(), newUser.Login, newUser.Password)

			if err == nil {
				claims, _ := token.ParseJWT(jwt)
				assert.Equal(t, test.expectedResult, claims.UserID.String())
			}

			assert.Equal(t, test.expectedErr, err)
		})
	}
}
