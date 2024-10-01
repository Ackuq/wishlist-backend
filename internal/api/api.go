package api

import (
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/api/handlers"
	"github.com/ackuq/wishlist-backend/internal/api/parser"
	"github.com/ackuq/wishlist-backend/internal/api/routes"
	"github.com/ackuq/wishlist-backend/internal/config"
	"github.com/ackuq/wishlist-backend/internal/db"
)

func New(queries *db.Queries, config *config.Config) error {
	parser := parser.New()
	handlers := handlers.New(queries, parser)

	router := routes.New(handlers)

	return http.ListenAndServe(config.Host, router)
}
