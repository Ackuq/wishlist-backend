package httperrors

import (
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/logger"
	"github.com/ackuq/wishlist-backend/internal/utils"
	"go.uber.org/zap"
)

func ErrorResponse(res http.ResponseWriter, req *http.Request, status int, message interface{}) {
	data := utils.JsonObject{"error": message}

	err := utils.WriteJSONResponse(res, status, data, nil)
	if err != nil {
		logger.Logger.Error(req, zap.Error(err))
		res.WriteHeader(500)
	}
}

func InternalServerErrorResponse(res http.ResponseWriter, req *http.Request, err error) {
	ErrorResponse(res, req, http.StatusInternalServerError, err.Error())
}

func NotFoundResponse(res http.ResponseWriter, req *http.Request, err error) {
	ErrorResponse(res, req, http.StatusNotFound, err.Error())
}

func BadRequestResponse(res http.ResponseWriter, req *http.Request, err error) {
	ErrorResponse(res, req, http.StatusBadRequest, err.Error())
}
