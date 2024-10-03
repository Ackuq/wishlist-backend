package routes

import (
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/api/handlers"
)

const authRoutePrefix = "/auth"

func authRoutes(handlers *handlers.Handlers) http.Handler {
	authRouter := http.NewServeMux()

	authRouter.Handle("GET /login", wrapHandler(handlers.AuthLogin))
	authRouter.Handle("GET /logout", wrapHandler(handlers.AuthLogout))
	authRouter.Handle("GET /callback", wrapHandler(handlers.AuthCallback))
	authRouter.Handle("GET /user", wrapHandler(handlers.AuthUser))

	return http.StripPrefix(authRoutePrefix, authRouter)
}
