package orders

import (
	"context"
	"github.com/VyacheslavKuzharov/gophermart/internal/entity"
	logdiscard "github.com/VyacheslavKuzharov/gophermart/internal/lib/log_discard"
	"github.com/VyacheslavKuzharov/gophermart/internal/repository"
	"github.com/VyacheslavKuzharov/gophermart/internal/repository/order"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUseCase_Upload(t *testing.T) {
	log := logdiscard.NewDiscardLogger()
	var pgUniqueFieldErr *repository.UniqueFieldErr

	testCases := map[string]struct {
		repo        order.RepositoryInterface
		expectedErr error
		orderNum    string
	}{
		"Happy path": {
			repo: &orderRepoMock{
				UploadError: nil,
			},
			orderNum:    "4026843483168683",
			expectedErr: nil,
		},
		"Order already exists": {
			repo: &orderRepoMock{
				UploadError: pgUniqueFieldErr,
			},
			expectedErr: pgUniqueFieldErr,
		},
	}

	for CaseName, test := range testCases {
		t.Run(CaseName, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), entity.CurrentUserID, uuid.NewV4())

			usecase := NewUseCase(test.repo, log)
			err := usecase.Upload(ctx, test.orderNum)

			assert.Equal(t, test.expectedErr, err)
		})
	}
}
