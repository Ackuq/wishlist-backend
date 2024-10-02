package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB = pgxpool.Pool

func New(ctx context.Context, connectionString string) (*DB, func(), error) {
	conn, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		return nil, nil, err
	}
	return conn, conn.Close, nil
}
