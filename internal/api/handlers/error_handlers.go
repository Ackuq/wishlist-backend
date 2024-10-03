package handlers

import (
	"log/slog"
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/api/models"
	"github.com/ackuq/wishlist-backend/internal/api/schemavalidator"
	"github.com/ackuq/wishlist-backend/internal/logger"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func HandleError(res http.ResponseWriter, req *http.Request, err error) {
	locale := req.Header.Get("Accept-Language")
	status, errors := errorToHttpObjects(err, locale)

	writeJSONResponse(res, status, models.ErrorResponse{Errors: errors})
}

func HandleCustomError(res http.ResponseWriter, err models.ErrorObject) {
	writeJSONResponse(res, err.Status, models.ErrorResponse{Errors: []models.ErrorObject{err}})
}

func errorToHttpObjects(err error, locale string) (int, []models.ErrorObject) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		return http.StatusBadRequest, schemavalidator.GetTranslationErrors(validationErrors, locale)
	}

	switch err {
	case pgx.ErrNoRows:
		return http.StatusNotFound, []models.ErrorObject{models.NotFoundError(err.Error())}
	}

	if dbError, ok := err.(*pgconn.PgError); ok {
		// Unique constraint failure
		if dbError.Code == "23505" {
			return http.StatusConflict,
				[]models.ErrorObject{models.ConflictError(err.Error())}
		}
	}
	slog.Error("Unknown server error", logger.ErrorAtr(err))
	return http.StatusInternalServerError,
		[]models.ErrorObject{models.ServerError(err.Error())}
}
