package orders

import (
	"context"
	"github.com/VyacheslavKuzharov/gophermart/internal/entity"
)

type orderRepoMock struct {
	UploadError error

	GetByNumberResult entity.Order
	GetByNumberErr    error

	GetManyByStatusesResult []entity.Order
	GetManyByStatusesErr    error

	UpdateStatusError error
}

func (repo *orderRepoMock) Create(ctx context.Context, orderDTO *entity.OrderDTO) error {
	return repo.UploadError
}

func (repo *orderRepoMock) GetByNumber(ctx context.Context, number string) (entity.Order, error) {
	return repo.GetByNumberResult, repo.GetByNumberErr
}

func (repo *orderRepoMock) GetManyByStatuses(ctx context.Context, statuses []string) ([]entity.Order, error) {
	return repo.GetManyByStatusesResult, repo.GetManyByStatusesErr
}

func (repo *orderRepoMock) UpdateStatus(ctx context.Context, status, number string) error {
	return repo.UpdateStatusError
}
