package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ackuq/wishlist-backend/internal/api/parser"
	"github.com/ackuq/wishlist-backend/internal/db"
	"github.com/ackuq/wishlist-backend/internal/logger"
	"go.uber.org/zap"
)

type Handlers struct {
	queries *db.Queries
	parser  *parser.Parser
}

func New(queries *db.Queries, parser *parser.Parser) *Handlers {
	return &Handlers{queries, parser}
}

func writeJSONResponse(res http.ResponseWriter, status int, data interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	js, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		logger.Logger.Error("Error marshaling JSON", zap.Error(err))
		return
	}

	_, err = res.Write(js)
	if err != nil {
		logger.Logger.Error("Error writing JSON response", zap.Error(err))
	}
}
