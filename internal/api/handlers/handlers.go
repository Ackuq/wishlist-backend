package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/api/schemavalidator"
	"github.com/ackuq/wishlist-backend/internal/db/queries"
	"github.com/ackuq/wishlist-backend/internal/logger"
)

type Handlers struct {
	queries         *queries.Queries
	schemaValidator *schemavalidator.SchemaValidator
}

func New(queries *queries.Queries, schemaValidator *schemavalidator.SchemaValidator) *Handlers {
	return &Handlers{queries, schemaValidator}
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
