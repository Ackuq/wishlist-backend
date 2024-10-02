package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/ackuq/wishlist-backend/internal/api"
	"github.com/ackuq/wishlist-backend/internal/config"
	"github.com/ackuq/wishlist-backend/internal/db"
	"github.com/ackuq/wishlist-backend/internal/logger"
	"github.com/jackc/pgx/v5"

	"github.com/joho/godotenv"
)

func main() {
	logger.InitLogger()

	err := godotenv.Load()
	if err != nil {
		slog.Error("Error reading .env", slog.Any("error", err))
		os.Exit(1)
	}

	config := config.GetConfig()

	// Connect to Postgres database
	conn, err := pgx.Connect(context.Background(), config.DataBase.URL)
	if err != nil {
		slog.Error("Error when connecting to DB", logger.ErrorAtr(err))
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// Init queries
	queries := db.New(conn)

	// Start API
	err = api.New(queries, config)
	if err != nil {
		slog.Error("Error starting API", logger.ErrorAtr(err))
		os.Exit(1)
	}
}
