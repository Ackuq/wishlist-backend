package repositories

import (
	"github.com/ackuq/wishlist-backend/internal/db"
)

type Repositories struct {
	UserRepository *UserRepository
}

func NewRepositories(db *db.Database) *Repositories {
	return &Repositories{
		UserRepository: NewUserRepository(db),
	}
}
