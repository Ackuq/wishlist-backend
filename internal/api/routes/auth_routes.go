package routes

import (
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/api/handlers"
	"github.com/ackuq/wishlist-backend/internal/api/middlewares"
)

const authRoutePrefix = "/auth"

func authRoutes(handlers *handlers.Handlers) http.Handler {
	authRouter := http.NewServeMux()

	authRouter.Handle("GET /login", http.HandlerFunc((handlers.AuthLogin)))
	authRouter.Handle("GET /logout", http.HandlerFunc(handlers.AuthLogout))
	authRouter.Handle("GET /callback", http.HandlerFunc(handlers.AuthCallback))
	authRouter.Handle("GET /user", middlewares.WithAuthentication(handlers.AuthUser))

	return http.StripPrefix(authRoutePrefix, authRouter)
}
