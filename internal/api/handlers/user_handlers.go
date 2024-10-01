package handlers

import (
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/api/models"
)

func (handlers *Handlers) CreateUser(res http.ResponseWriter, req *http.Request) {
	var createUser models.CreateUser

	body := &models.CreateUser{}
	if err := handlers.parser.BindJSON(req, body); err != nil {
		handleError(res, req, err)
		return
	}

	err := handlers.queries.CreateUser(&createUser)

	if err != nil {
		handleError(res, req, err)
		return
	}

	res.WriteHeader(http.StatusCreated)
}
