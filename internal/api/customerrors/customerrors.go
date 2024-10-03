package customerrors

import (
	"github.com/ackuq/wishlist-backend/internal/api/models"
)

// JSON errors
var JSONDecodingError = models.BadRequestError("Could not decode json into desired type.")
var InvalidResultTypeError = models.ServerError("Invalid result type.")

// Auth errors
var InvalidReturnToURL = models.BadRequestError("Invalid `return_to` URL.")
var InvalidStateParameterError = models.BadRequestError("Invalid state parameter.")
var InvalidSessionStateError = models.BadRequestError("Invalid session state.")
var ExchangeFailedError = models.UnauthorizedError("Failed to exchange an authorization code for a token.")
var Unauthenticated = models.UnauthorizedError("Authentication not found.")
var VerifyFailedError = models.ServerError("Failed to verify ID Token")
