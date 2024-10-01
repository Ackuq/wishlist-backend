package routes

import (
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/api/handlers"
)

const userRoutePrefix = "/api/v1/user"

func userRoutes(handlers *handlers.Handlers) http.Handler {
	userRouter := http.NewServeMux()

	userRouter.Handle("POST /", wrapHandler(handlers.CreateUser))

	return http.StripPrefix(userRoutePrefix, userRouter)
}
