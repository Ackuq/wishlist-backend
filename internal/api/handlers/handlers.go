package handlers

import "github.com/ackuq/wishlist-backend/internal/db/repositories"

type Handlers struct {
	UserHandler *UserHandler
}

func NewHandlers(repo *repositories.Repositories) *Handlers {
	return &Handlers{
		UserHandler: NewUserHandler(repo.UserRepository),
	}
}
