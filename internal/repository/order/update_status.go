package order

import (
	"context"
)

func (repo *Repository) UpdateStatus(ctx context.Context, status, number string) error {
	log := repo.logger.Logger

	log.Info().Msgf("RUN SQL: UPDATE orders SET status = %s WHERE number = %s", status, number)
	_, err := repo.pg.Pool.Exec(
		ctx,
		`UPDATE orders SET status = $1 WHERE number = $2`,
		status,
		number,
	)
	if err != nil {
		return err
	}

	return nil
}
