package main

import (
	"github.com/ackuq/wishlist-backend/internal/api"
	"github.com/ackuq/wishlist-backend/internal/api/handlers"
	"github.com/ackuq/wishlist-backend/internal/config"
	"github.com/ackuq/wishlist-backend/internal/db"
	"github.com/ackuq/wishlist-backend/internal/db/repositories"

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

	db := db.Connect(config)
	defer db.Close()

	repositories := repositories.NewRepositories(db)
	handlers := handlers.NewHandlers(repositories)

	err = api.NewApi(handlers, config)
	if err != nil {
		logger.Logger.Fatal(err)
	}
}
