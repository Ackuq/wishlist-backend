package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/api/httperrors"
	"github.com/ackuq/wishlist-backend/internal/api/models"
	"github.com/ackuq/wishlist-backend/internal/db/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	UserRepository *repositories.UserRepository
}

func NewUserHandler(ur *repositories.UserRepository) *UserHandler {
	return &UserHandler{
		UserRepository: ur,
	}
}

func (uh *UserHandler) CreateUser() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var createUser models.CreateUser

		decoder := json.NewDecoder(req.Body)

		err := decoder.Decode(&createUser)

		if err != nil {
			httperrors.BadRequestResponse(res, req, errors.New("invalid request payload"))
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword(createUser.Password, 14)

		if err != nil {
			httperrors.InternalServerErrorResponse(res, req, errors.New("password hashing failed"))
			return
		}

		createUser.Password = hashedPassword

		err = uh.UserRepository.CreateUser(&createUser)

		if err != nil {
			httperrors.InternalServerErrorResponse(res, req, errors.New("failed to create user"))
			return
		}

		res.WriteHeader(http.StatusCreated)
	})
}
