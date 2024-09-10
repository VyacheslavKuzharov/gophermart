package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
	"os"
	"time"
)

type JWTClaim struct {
	jwt.RegisteredClaims
	UserID uuid.UUID
}

func CreateJWT(userID uuid.UUID) (string, error) {
	key := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * 72)},
		},
		UserID: userID,
	})

	tokenStr, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ParseJWT(jwtToken string) (*JWTClaim, error) {
	key := []byte(os.Getenv("JWT_SECRET"))
	var claims JWTClaim

	token, err := jwt.ParseWithClaims(jwtToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return &claims, nil
}
