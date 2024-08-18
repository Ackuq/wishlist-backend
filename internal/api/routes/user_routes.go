package routes

import (
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/api/handlers"
)

const UserRoutePrefix = "/api/v1/user"

func UserMux(handlers *handlers.Handlers) http.Handler {
	userMux := http.NewServeMux()

	userMux.Handle("POST /", handlers.UserHandler.CreateUser())

	return http.StripPrefix(UserRoutePrefix, userMux)
}
