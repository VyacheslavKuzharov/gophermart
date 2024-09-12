package middlewares

import (
	"context"
	"errors"
	"github.com/VyacheslavKuzharov/gophermart/internal/entity"
	"github.com/VyacheslavKuzharov/gophermart/internal/lib/response"
	"github.com/VyacheslavKuzharov/gophermart/internal/lib/token"
	"github.com/VyacheslavKuzharov/gophermart/internal/transport/http/handlers/auth"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		jwt, err := fetchToken(r, auth.CookieName)
		if err != nil {
			response.Err(w, err.Error(), http.StatusUnauthorized)
			return
		}

		claims, err := token.ParseJWT(jwt)
		if err != nil {
			response.Err(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), entity.CurrentUserID, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

func fetchToken(r *http.Request, name string) (string, error) {
	authHeader := r.Header.Get(name)
	if authHeader != "" {
		return authHeader, nil
	}

	cookie, err := r.Cookie(name)
	if err != nil {
		return "", err
	}
	if cookie.Value == "" {
		return "", errors.New("token blank")
	}

	return cookie.Value, nil
}
