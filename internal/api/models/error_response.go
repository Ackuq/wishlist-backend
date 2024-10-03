package models

import "net/http"

type ErrorResponse struct {
	Errors []ErrorObject `json:"errors"`
}

type ErrorObject struct {
	Type    string `json:"type"`
	Status  int    `json:"code"`
	Message string `json:"message"`
}

func ValidationError(message string) ErrorObject {
	return ErrorObject{
		Type:    "ValidationError",
		Status:  http.StatusBadRequest,
		Message: message,
	}
}

func ServerError(message string) ErrorObject {
	return ErrorObject{
		Type:    "ServerError",
		Status:  http.StatusInternalServerError,
		Message: message,
	}
}

func ConflictError(message string) ErrorObject {
	return ErrorObject{
		Type:    "ConflictError",
		Status:  http.StatusConflict,
		Message: message,
	}
}

func NotFoundError(message string) ErrorObject {
	return ErrorObject{
		Type:    "NotFoundError",
		Status:  http.StatusNotFound,
		Message: message,
	}
}

func BadRequestError(message string) ErrorObject {
	return ErrorObject{
		Type:    "BadRequest",
		Status:  http.StatusBadRequest,
		Message: message,
	}
}

func UnauthorizedError(message string) ErrorObject {
	return ErrorObject{
		Type:    "Unauthorized",
		Status:  http.StatusUnauthorized,
		Message: message,
	}
}
