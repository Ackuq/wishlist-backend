package api

import (
	"fmt"
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/api/handlers"
	"github.com/ackuq/wishlist-backend/internal/api/routes"
	"github.com/ackuq/wishlist-backend/internal/api/schemavalidator"
	"github.com/ackuq/wishlist-backend/internal/config"
	"github.com/ackuq/wishlist-backend/internal/db"
	"github.com/ackuq/wishlist-backend/internal/logger"
)

func New(queries *db.Queries, config *config.Config) error {
	schemaValidator := schemavalidator.New()
	handlers := handlers.New(queries, schemaValidator)

	router := routes.New(handlers)

	logger.Logger.Info(fmt.Sprintf("Listing on host %s", config.Host))
	return http.ListenAndServe(config.Host, router)
}
