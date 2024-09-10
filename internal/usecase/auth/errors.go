package auth

import (
	"fmt"
)

type ErrGenPwdHash struct {
	Value string
	Err   error
}

func (e *ErrGenPwdHash) Error() string {
	return fmt.Sprintf("could not generate password hash from value: %s. %v", e.Value, e.Err)
}

func NewErrGenPwdHash(v string, err error) error {
	return &ErrGenPwdHash{
		Value: v,
		Err:   err,
	}
}
