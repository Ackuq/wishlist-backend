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
	"github.com/rs/cors"
)

func New(queries *queries.Queries, config *config.Config) error {
	err := auth.Init(config)
	if err != nil {
		slog.Error("Error setting up auth", logger.ErrorAtr(err))
		os.Exit(1)
	}

	sessionmanager.Init()
	schemavalidator.Init()
	handlers := handlers.New(queries)

	router := routes.New(handlers)

	cors := cors.New(cors.Options{
		AllowedOrigins: config.CORS.AllowedOrigins,
		AllowedMethods: []string{http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete},
		AllowCredentials: true,
	})

	slog.Info(fmt.Sprintf("Listing on host %s", config.Host))

	return http.ListenAndServe(config.Host, withMiddlewares(router, cors))
}

func withMiddlewares(router *http.ServeMux, cors *cors.Cors) http.Handler {
	// LoadAndSave handles loading and committing session data to the session store
	withSessionManager := sessionmanager.Get().LoadAndSave(router)
	// TODO: This should be more strict if this ever gets deployed...
	withCors := cors.Handler(withSessionManager)

	return withCors
}
