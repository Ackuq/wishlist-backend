package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	pgxuuid "github.com/vgarvardt/pgx-google-uuid/v5"
)

type DB = pgxpool.Pool

func New(ctx context.Context, connectionString string) (*DB, func(), error) {
	dbConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, nil, err
	}

	dbConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxuuid.Register(conn.TypeMap())
		return nil
	}

	conn, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		return nil, nil, err
	}

	return conn, conn.Close, nil
}
