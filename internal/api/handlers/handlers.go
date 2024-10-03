package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"reflect"

	"github.com/ackuq/wishlist-backend/internal/api/auth"
	"github.com/ackuq/wishlist-backend/internal/api/customerrors"
	"github.com/ackuq/wishlist-backend/internal/api/schemavalidator"
	"github.com/ackuq/wishlist-backend/internal/db/queries"
	"github.com/ackuq/wishlist-backend/internal/logger"
	"github.com/alexedwards/scs/v2"
)

type Handlers struct {
	queries         *queries.Queries
	schemaValidator *schemavalidator.SchemaValidator
	auth            *auth.Authenticator
	sessionManager  *scs.SessionManager
}

func New(queries *queries.Queries, schemaValidator *schemavalidator.SchemaValidator, auth *auth.Authenticator, sessionManager *scs.SessionManager) *Handlers {
	return &Handlers{queries, schemaValidator, auth, sessionManager}
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
		handlers.handleCustomError(res, customerrors.JSONDecodingError)
		return false
	}

	value := reflect.ValueOf(result)

	switch value.Kind() {
	case reflect.Ptr:
		if err := handlers.schemaValidator.Struct(value.Elem().Interface()); err != nil {
			handlers.handleError(res, req, err)
			return false
		}
		return true
	case reflect.Struct:
		if err := handlers.schemaValidator.Struct(result); err != nil {
			handlers.handleError(res, req, err)
			return false
		}
		return true
	}

	handlers.handleCustomError(res, customerrors.InvalidResultTypeError)
	return false
}
