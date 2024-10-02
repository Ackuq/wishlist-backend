package routes

import (
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/api/handlers"
)

func New(handlers *handlers.Handlers) *http.ServeMux {
	router := http.NewServeMux()

	router.Handle(accountRoutePrefix+"/", accountRoutes(handlers))

	return router
}

func wrapHandler(handler func(res http.ResponseWriter, req *http.Request)) http.Handler {
	return http.HandlerFunc(handler)
}
