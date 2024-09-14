package orders

import "context"

type ordersUseCaseMock struct {
	UploadError error
}

func (uc *ordersUseCaseMock) Upload(ctx context.Context, orderNumber string) error {
	return uc.UploadError
}
