package utils

import (
	"encoding/json"
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/logger"
	"go.uber.org/zap"
)

type JsonObject map[string]interface{}

func WriteJSONResponse(res http.ResponseWriter, status int, data JsonObject, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		logger.Logger.Error("Error marshaling JSON", zap.Error(err))
		return err
	}

	for key, value := range headers {
		res.Header()[key] = value
		return err
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)

	_, err = res.Write(js)
	if err != nil {
		logger.Logger.Error("Error writing JSON response", zap.Error(err))
		return err
	}
	return nil
}
