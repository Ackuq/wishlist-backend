package repositories

import (
	"github.com/ackuq/wishlist-backend/internal/api/models"
	"github.com/ackuq/wishlist-backend/internal/db"
)

type UserRepository struct {
	db *db.Database
}

func NewUserRepository(db *db.Database) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(user *models.CreateUser) error {
	// TODO
	return nil
}
