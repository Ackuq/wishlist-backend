package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ackuq/wishlist-backend/internal/api/auth"
	"github.com/ackuq/wishlist-backend/internal/api/customerrors"
	"github.com/ackuq/wishlist-backend/internal/api/models"
	"github.com/ackuq/wishlist-backend/internal/api/sessionmanager"
)

func (handlers *Handlers) AuthLogin(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	state, err := generateInitialState()

	if err != nil {
		HandleError(res, req, err)
		return
	}

	// Save state inside session
	sessionManager := sessionmanager.Get()
	sessionManager.Put(ctx, auth.StateSessionKey, state)

	http.Redirect(res, req, handlers.auth.AuthCodeURL(state), http.StatusTemporaryRedirect)
}

func generateInitialState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}

func (handlers *Handlers) AuthCallback(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	sessionManager := sessionmanager.Get()

	sessionState := sessionManager.GetString(ctx, auth.StateSessionKey)
	queryParams := req.URL.Query()
	// Verify state is valid
	if sessionState != queryParams.Get("state") {
		HandleCustomError(res, customerrors.InvalidStateParameterError)
		return
	}

	token, err := handlers.auth.Exchange(ctx, queryParams.Get("code"))
	if err != nil {
		HandleCustomError(res, customerrors.ExchangeFailedError)
		return
	}

	idToken, err := handlers.auth.VerifyIDToken(ctx, token)
	if err != nil {
		HandleCustomError(res, customerrors.VerifyFailedError)
		return
	}

	var claims auth.Claims
	if err := idToken.Claims(&claims); err != nil {
		HandleError(res, req, err)
		return
	}

	sessionManager.Put(ctx, auth.AccessTokenSessionKey, token.AccessToken)
	sessionManager.Put(ctx, auth.ClaimsSessionKey, claims)

	res.WriteHeader(http.StatusCreated)
}

func (handlers *Handlers) AuthLogout(res http.ResponseWriter, req *http.Request) {
	sessionManager := sessionmanager.Get()
	logoutUrl, err := url.Parse(handlers.auth.LogoutUrl)
	if err != nil {
		HandleError(res, req, err)
		return
	}

	scheme := "http"
	if req.TLS != nil {
		scheme = "https"
	}

	returnTo, err := url.Parse(fmt.Sprintf("%s://%s", scheme, req.Host))
	if err != nil {
		HandleError(res, req, err)
		return
	}

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", handlers.auth.ClientID)
	logoutUrl.RawQuery = parameters.Encode()

	// Destroy session
	sessionManager.Destroy(req.Context())

	// Unauthenticate with Auth0
	http.Redirect(res, req, logoutUrl.String(), http.StatusTemporaryRedirect)
}

func (handlers *Handlers) AuthUser(res http.ResponseWriter, req *http.Request, claims auth.Claims) {
	user := models.User{
		Name:  claims.Name,
		Email: claims.Email,
	}
	writeJSONResponse(res, http.StatusOK, user)
}
