package token

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateJWT(t *testing.T) {
	ID := uuid.NewV4()

	testCases := map[string]struct {
		id             uuid.UUID
		expectedResult uuid.UUID
		expectedErr    error
	}{
		"Creates jwt token": {
			id:             ID,
			expectedResult: ID,
			expectedErr:    nil,
		},
	}

	for CaseName, test := range testCases {
		t.Run(CaseName, func(t *testing.T) {
			jwt, err := CreateJWT(ID)
			claims, _ := ParseJWT(jwt)
			assert.Equal(t, test.expectedResult, claims.UserID)

			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestParseJWT(t *testing.T) {
	ID, _ := uuid.FromString("ff8bb8f1-e0d3-4819-a4da-c140919128f3")
	jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJmZjhiYjhmMS1lMGQzLTQ4MTktYTRkYS1jMTQwOTE5MTI4ZjMifQ.Ai7AbEhnfFTgglJbGm02N9u66tlhANYstjLLmsUHurg"
	invalitJWT := "invaid"

	testCases := map[string]struct {
		jwt            string
		expectedResult uuid.UUID
		expectedErr    error
	}{
		"Parse jwt token": {
			jwt:            jwt,
			expectedResult: ID,
			expectedErr:    nil,
		},
		"Error": {
			jwt:            invalitJWT,
			expectedResult: ID,
			expectedErr:    errors.New("token is malformed"),
		},
	}

	for CaseName, test := range testCases {
		t.Run(CaseName, func(t *testing.T) {
			claims, err := ParseJWT(test.jwt)
			if err == nil {
				assert.Equal(t, test.expectedResult, claims.UserID)
			}

			assert.Equal(t, test.expectedErr, errors.Unwrap(err))
		})
	}
}
