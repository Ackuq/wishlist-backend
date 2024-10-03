package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/ackuq/wishlist-backend/internal/api/auth"
	"github.com/ackuq/wishlist-backend/internal/api/handlers"
	"github.com/ackuq/wishlist-backend/internal/api/routes"
	"github.com/ackuq/wishlist-backend/internal/api/schemavalidator"
	"github.com/ackuq/wishlist-backend/internal/api/sessionmanager"
	"github.com/ackuq/wishlist-backend/internal/config"
	"github.com/ackuq/wishlist-backend/internal/db/queries"
	"github.com/ackuq/wishlist-backend/internal/logger"
)

func New(queries *queries.Queries, config *config.Config) error {
	auth, err := auth.New(config)
	if err != nil {
		slog.Error("Error setting up auth", logger.ErrorAtr(err))
		os.Exit(1)
	}

	sessionManager := sessionmanager.New()
	schemaValidator := schemavalidator.New()
	handlers := handlers.New(queries, schemaValidator, auth, sessionManager)

	router := routes.New(handlers)

	slog.Info(fmt.Sprintf("Listing on host %s", config.Host))
	// LoadAndSave handles loading and committing session data to the session store
	return http.ListenAndServe(config.Host, sessionManager.LoadAndSave(router))
}
