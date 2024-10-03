package customerrors

import (
	"github.com/ackuq/wishlist-backend/internal/api/models"
)

// JSON errors
var JSONDecodingError = models.BadRequestError("Could not decode json into desired type.")
var InvalidResultTypeError = models.ServerError("Invalid result type.")

// Auth errors
var InvalidStateParameterError = models.BadRequestError("Invalid state parameter.")
var ExchangeFailedError = models.UnauthorizedError("Failed to exchange an authorization code for a token.")
var VerifyFailedError = models.ServerError("Failed to verify ID Token")
