package handlers

import (
	"fmt"
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/api/models"
	"github.com/ackuq/wishlist-backend/internal/logger"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"
)

func handleError(res http.ResponseWriter, req *http.Request, err error) {
	logger.Logger.Error(req, zap.Error(err))

	status, errors := errorToHttpObjects(err)

	writeJSONResponse(res, status, models.ErrorResponse{Errors: errors})
}

func errorToHttpObjects(err error) (int, []models.ErrorObject) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		return http.StatusBadRequest, getValidationErrorObjects(validationErrors)
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

	return http.StatusInternalServerError,
		[]models.ErrorObject{models.ServerError(err.Error())}
}

func getValidationErrorObjects(errors validator.ValidationErrors) []models.ErrorObject {
	errorObjects := make([]models.ErrorObject, len(errors))
	for i, fieldError := range errors {
		message := GetValidationErrorMessage(fieldError)
		errorObjects[i] = models.ValidationError(message)
	}
	return errorObjects
}

func GetValidationErrorMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fieldError.Field())
	case "lte":
		return fmt.Sprintf("%s should be less than or equal to %s", fieldError.Field(), fieldError.Param())
	case "lt":
		return fmt.Sprintf("%s should be less than %s", fieldError.Field(), fieldError.Param())
	case "gte":
		return fmt.Sprintf("%s should be greater than or equal to %s", fieldError.Field(), fieldError.Param())
	case "gt":
		return fmt.Sprintf("%s should be greater than %s", fieldError.Field(), fieldError.Param())
	case "min":
		return fmt.Sprintf("%s should have minimum length of %s", fieldError.Field(), fieldError.Param())
	case "max":
		return fmt.Sprintf("%s should have maximum length of %s", fieldError.Field(), fieldError.Param())
	case "alpha":
		return fmt.Sprintf("%s should contain alpha characters only", fieldError.Field())
	}
	return "Unknown error"
}
