package routes

import (
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/api/handlers"
)

func New(handlers *handlers.Handlers) *http.ServeMux {
	router := http.NewServeMux()

	router.Handle(accountRoutePrefix+"/", accountRoutes(handlers))
	router.Handle(authRoutePrefix+"/", authRoutes(handlers))

	return router
}
