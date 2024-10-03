package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ackuq/wishlist-backend/internal/api/customerrors"
)

func (handlers *Handlers) AuthLogin(res http.ResponseWriter, req *http.Request) {
	state, err := generateInitialState()

	if err != nil {
		handlers.handleError(res, req, err)
		return
	}

	// Save state inside session
	handlers.sessionManager.Put(req.Context(), "state", state)

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
	sessionState := handlers.sessionManager.GetString(req.Context(), "state")
	queryParams := req.URL.Query()
	// Verify state is valid
	if sessionState != queryParams.Get("state") {
		handlers.handleCustomError(res, customerrors.InvalidStateParameterError)
		return
	}

	token, err := handlers.auth.Exchange(req.Context(), queryParams.Get("code"))
	if err != nil {
		handlers.handleCustomError(res, customerrors.ExchangeFailedError)
		return
	}

	idToken, err := handlers.auth.VerifyIDToken(req.Context(), token)
	if err != nil {
		handlers.handleCustomError(res, customerrors.VerifyFailedError)
		return
	}

	var profile map[string]any
	if err := idToken.Claims(&profile); err != nil {
		handlers.handleError(res, req, err)
		return
	}

	handlers.sessionManager.Put(req.Context(), "access_token", token.AccessToken)
	handlers.sessionManager.Put(req.Context(), "profile", profile)

	res.WriteHeader(http.StatusCreated)
}

func (handlers *Handlers) AuthLogout(res http.ResponseWriter, req *http.Request) {
	logoutUrl, err := url.Parse(handlers.auth.LogoutUrl)
	if err != nil {
		handlers.handleError(res, req, err)
		return
	}

	scheme := "http"
	if req.TLS != nil {
		scheme = "https"
	}

	returnTo, err := url.Parse(fmt.Sprintf("%s://%s", scheme, req.Host))
	if err != nil {
		handlers.handleError(res, req, err)
		return
	}
	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", handlers.auth.ClientID)
	logoutUrl.RawQuery = parameters.Encode()

	http.Redirect(res, req, logoutUrl.String(), http.StatusTemporaryRedirect)
}
