package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"reflect"

	"github.com/ackuq/wishlist-backend/internal/api/customerrors"
	"github.com/ackuq/wishlist-backend/internal/api/schemavalidator"
	"github.com/ackuq/wishlist-backend/internal/db/queries"
	"github.com/ackuq/wishlist-backend/internal/logger"
)

type Handlers struct {
	queries *queries.Queries
}

func New(queries *queries.Queries) *Handlers {
	return &Handlers{queries}
}

func writeJSONResponse(res http.ResponseWriter, status int, data interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	js, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		slog.Error("Error marshaling JSON", logger.ErrorAtr(err))
		return
	}

	_, err = res.Write(js)
	if err != nil {
		slog.Error("Error writing JSON response", logger.ErrorAtr(err))
	}
}
func (handlers *Handlers) bindJSON(res http.ResponseWriter, req *http.Request, result any) bool {
	err := json.NewDecoder(req.Body).Decode(result)

	if err != nil {
		HandleCustomError(res, customerrors.JSONDecodingError)
		return false
	}

	value := reflect.ValueOf(result)

	switch value.Kind() {
	case reflect.Ptr:
		if err := schemavalidator.ValidateStruct(value.Elem().Interface()); err != nil {
			HandleError(res, req, err)
			return false
		}
		return true
	case reflect.Struct:
		if err := schemavalidator.ValidateStruct(result); err != nil {
			HandleError(res, req, err)
			return false
		}
		return true
	}

	HandleCustomError(res, customerrors.InvalidResultTypeError)
	return false
}
