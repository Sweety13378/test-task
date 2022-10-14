package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Conn *pgxpool.Pool

func Connect(ctx context.Context, config Config) error {
	var err error
	for i := 0; i != config.MaxAttempts; i++ {
		Conn, err = pgxpool.New(ctx, config.DsnString())
		if err != nil {
			continue
		}
		err = Conn.Ping(ctx)
		if err != nil {
			return nil
		}
		break
	}
	return err
}
