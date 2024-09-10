package auth

import (
	"context"
)

type authUseCaseMock struct {
	SignUpResult string
	SignUpError  error

	SignInResult string
	SignInError  error
}

func (uc *authUseCaseMock) SignUp(ctx context.Context, login, password string) (string, error) {
	return uc.SignUpResult, uc.SignUpError
}

func (uc *authUseCaseMock) SignIn(ctx context.Context, login, password string) (string, error) {
	return uc.SignInResult, uc.SignInError
}
