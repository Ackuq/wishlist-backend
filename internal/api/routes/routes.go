package routes

import (
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/api/handlers"
)

func InitializeRoutes(handlers *handlers.Handlers) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle(UserRoutePrefix+"/", UserMux(handlers))

	return mux
}
