package models

type ErrorResponse struct {
	Errors []ErrorObject `json:"errors"`
}

type ErrorObject struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func ValidationError(message string) ErrorObject {
	return ErrorObject{
		Type:    "ValidationError",
		Message: message,
	}
}

func ServerError(message string) ErrorObject {
	return ErrorObject{
		Type:    "ServerError",
		Message: message,
	}
}

func ConflictError(message string) ErrorObject {
	return ErrorObject{
		Type:    "ConflictError",
		Message: message,
	}
}

func NotFoundError(message string) ErrorObject {
	return ErrorObject{
		Type:    "NotFoundError",
		Message: message,
	}
}
