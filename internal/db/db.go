package db

import (
	"context"

	"github.com/ackuq/wishlist-backend/internal/config"
	"github.com/ackuq/wishlist-backend/internal/logger"
	"github.com/jackc/pgx/v5"
)

type Database struct {
	conn pgx.Conn
}

func Connect(config *config.Config) *Database {
	conn, err := pgx.Connect(context.Background(), config.DataBase.URL)
	if err != nil {
		logger.Logger.Fatal(err)
	}
	return &Database{
		conn: *conn,
	}
}

func (db *Database) Close() {
	db.conn.Close(context.Background())
}
