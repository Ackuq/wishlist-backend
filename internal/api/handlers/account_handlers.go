package handlers

import (
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/api/models"
	"github.com/ackuq/wishlist-backend/internal/db/queries"
	"github.com/google/uuid"
)

func (handlers *Handlers) CreateAccount(res http.ResponseWriter, req *http.Request) {
	body := &models.CreateAccount{}
	if ok := handlers.bindJSON(res, req, body); !ok {
		return
	}

	_, err := handlers.queries.CreateAccount(req.Context(), queries.CreateAccountParams{
		Name:  body.Name,
		Email: body.Email,
	})

	if err != nil {
		handlers.handleError(res, req, err)
		return
	}

	res.WriteHeader(http.StatusCreated)
}

func (handlers *Handlers) GetAccount(res http.ResponseWriter, req *http.Request) {
	id, err := uuid.Parse(req.PathValue("id"))

	if err != nil {
		handlers.handleError(res, req, err)
		return
	}

	account, err := handlers.queries.GetAccount(req.Context(), id)

	if err != nil {
		handlers.handleError(res, req, err)
		return
	}

	response := models.Account{
		ID:    account.ID,
		Name:  account.Name,
		Email: account.Email,
	}

	writeJSONResponse(res, http.StatusOK, response)
}

func (handlers *Handlers) ListAccounts(res http.ResponseWriter, req *http.Request) {
	accounts, err := handlers.queries.ListAccounts(req.Context())

	if err != nil {
		handlers.handleError(res, req, err)
	}

	response := make([]models.Account, len(accounts))
	for i, account := range accounts {
		response[i] = models.Account{
			ID:    account.ID,
			Name:  account.Name,
			Email: account.Email,
		}
	}

	writeJSONResponse(res, http.StatusOK, response)
}
