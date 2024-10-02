package routes

import (
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/api/handlers"
)

const accountRoutePrefix = "/api/v1/account"

func accountRoutes(handlers *handlers.Handlers) http.Handler {
	accountRouter := http.NewServeMux()

	accountRouter.Handle("POST /", wrapHandler(handlers.CreateAccount))
	accountRouter.Handle("GET /{id}", wrapHandler(handlers.GetAccount))
	accountRouter.Handle("GET /", wrapHandler(handlers.ListAccounts))

	return http.StripPrefix(accountRoutePrefix, accountRouter)
}
