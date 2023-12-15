package store

import (
	"context"
	"gameon-twotwentyk-api/connections"
)

func AddBalanceToUser(ctx context.Context, user_id int64, amount int64) error {
	q := `UPDATE public.user SET balance = balance + $1 WHERE id = $2`
	_, err := connections.Postgres.ExecContext(ctx, q, amount, user_id)
	if err != nil {
		return err
	}

	return err
}
