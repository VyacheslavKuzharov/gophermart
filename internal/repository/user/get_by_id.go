package user

import (
	"context"
	"github.com/VyacheslavKuzharov/gophermart/internal/entity"
)

func (r *Repository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	user := new(entity.User)

	//var originalURL string
	//
	//row := r.pg.Pool.QueryRow(
	//	ctx,
	//	"SELECT original_url FROM shorten_urls WHERE short_key = $1 AND is_deleted = false",
	//	key,
	//)
	//err := row.Scan(&originalURL)
	//if err != nil {
	//	return "", errors.New("shortKey not found")
	//}
	//
	//return originalURL, nil

	return user, nil
}
