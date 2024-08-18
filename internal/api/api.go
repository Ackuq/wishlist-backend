package api

import (
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/api/handlers"
	"github.com/ackuq/wishlist-backend/internal/api/routes"
	"github.com/ackuq/wishlist-backend/internal/config"
)

func NewApi(handlers *handlers.Handlers, config *config.Config) error {
	mux := routes.InitializeRoutes(handlers)

	return http.ListenAndServe(config.Host, mux)
}
