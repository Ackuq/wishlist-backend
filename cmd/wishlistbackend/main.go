package main

import (
	"context"

	"github.com/ackuq/wishlist-backend/internal/api"
	"github.com/ackuq/wishlist-backend/internal/config"
	"github.com/ackuq/wishlist-backend/internal/db"
	"github.com/jackc/pgx/v5"

	"github.com/ackuq/wishlist-backend/internal/logger"
	"github.com/joho/godotenv"
)

func main() {
	logger.InitLogger()
	defer logger.CloseLogger()

	err := godotenv.Load()
	if err != nil {
		logger.Logger.Fatal(err)
	}

	config := config.GetConfig()

	// Connect to Postgres database
	conn, err := pgx.Connect(context.Background(), config.DataBase.URL)
	if err != nil {
		logger.Logger.Fatal(err)
	}
	defer conn.Close(context.Background())

	// Init queries
	queries := db.New(conn)

	// Start API
	err = api.New(queries, config)
	if err != nil {
		logger.Logger.Fatal(err)
	}
}
