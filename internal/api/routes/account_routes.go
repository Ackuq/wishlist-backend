package routes

import (
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/api/handlers"
)

const accountRoutePrefix = "/api/v1/account"

func accountRoutes(handlers *handlers.Handlers) http.Handler {
	accountRouter := http.NewServeMux()

	accountRouter.Handle("POST /", http.HandlerFunc(handlers.CreateAccount))
	accountRouter.Handle("GET /{id}", http.HandlerFunc(handlers.GetAccount))
	accountRouter.Handle("GET /", http.HandlerFunc(handlers.ListAccounts))

	return http.StripPrefix(accountRoutePrefix, accountRouter)
}
