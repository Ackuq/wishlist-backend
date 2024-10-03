package middlewares

import (
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/api/auth"
	"github.com/ackuq/wishlist-backend/internal/api/customerrors"
	"github.com/ackuq/wishlist-backend/internal/api/handlers"
	"github.com/ackuq/wishlist-backend/internal/api/sessionmanager"
)

func WithAuthentication(next func(http.ResponseWriter, *http.Request, auth.Claims)) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		sessionManager := sessionmanager.Get()
		claims, ok := sessionManager.Get(ctx, auth.ClaimsSessionKey).(auth.Claims)

		if !ok {
			handlers.HandleCustomError(res, customerrors.Unauthenticated)
			return
		}

		next(res, req, claims)
	})
}
