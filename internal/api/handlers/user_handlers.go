package handlers

import (
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/api/models"
)

func (handlers *Handlers) CreateUser(res http.ResponseWriter, req *http.Request) {
	body := &models.CreateUser{}
	if err := handlers.schemaValidator.BindJSON(req, body); err != nil {
		handlers.handleError(res, req, err)
		return
	}

	err := handlers.queries.CreateUser(body)

	if err != nil {
		handlers.handleError(res, req, err)
		return
	}

	res.WriteHeader(http.StatusCreated)
}
