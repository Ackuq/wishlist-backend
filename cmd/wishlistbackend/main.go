package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/ackuq/wishlist-backend/internal/api"
	"github.com/ackuq/wishlist-backend/internal/config"
	"github.com/ackuq/wishlist-backend/internal/db"
	"github.com/ackuq/wishlist-backend/internal/db/queries"
	"github.com/ackuq/wishlist-backend/internal/logger"

	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	logger.InitLogger()

	// Read .env
	if err := godotenv.Load(); err != nil {
		slog.Error("Error reading .env", slog.Any("error", err))
		os.Exit(1)
	}
	config := config.GetConfig()

	// Migrate database
	if err := db.Migrate(config.Database.URL); err != nil {
		slog.Error("Error migrating DB", logger.ErrorAtr(err))
		os.Exit(1)
	}

	// Connect to database
	db, dbDispose, err := db.New(ctx, config.Database.URL)
	if err != nil {
		slog.Error("Error when connecting to DB", logger.ErrorAtr(err))
		os.Exit(1)
	}
	defer dbDispose()

	// Init queries
	queries := queries.New(db)

	// Start API
	err = api.New(queries, config)
	if err != nil {
		slog.Error("Error starting API", logger.ErrorAtr(err))
		os.Exit(1)
	}
}
