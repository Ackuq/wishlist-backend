package handlers

import (
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/api/models"
	"github.com/ackuq/wishlist-backend/internal/logger"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"
)

func (handlers *Handlers) handleError(res http.ResponseWriter, req *http.Request, err error) {

	locale := req.Header.Get("Accept-Language")
	status, errors := handlers.errorToHttpObjects(err, locale)

	writeJSONResponse(res, status, models.ErrorResponse{Errors: errors})
}

func (handlers *Handlers) errorToHttpObjects(err error, locale string) (int, []models.ErrorObject) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		return http.StatusBadRequest, handlers.schemaValidator.GetTranslationErrors(validationErrors, locale)
	}

	if err == pgx.ErrNoRows {
		return http.StatusNotFound, []models.ErrorObject{models.NotFoundError(err.Error())}
	}

	if dbError, ok := err.(*pgconn.PgError); ok {
		// Unique constraint failure
		if dbError.Code == "23505" {
			return http.StatusConflict,
				[]models.ErrorObject{models.ConflictError(err.Error())}
		}
	}
	logger.Logger.Error(zap.Error(err))
	return http.StatusInternalServerError,
		[]models.ErrorObject{models.ServerError(err.Error())}
}
