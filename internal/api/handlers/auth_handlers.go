package handlers

import (
	"net/http"
	"net/url"

	"github.com/ackuq/wishlist-backend/internal/api/auth"
	"github.com/ackuq/wishlist-backend/internal/api/customerrors"
	"github.com/ackuq/wishlist-backend/internal/api/models"
	"github.com/ackuq/wishlist-backend/internal/api/sessionmanager"
	"github.com/ackuq/wishlist-backend/internal/utils"
)

func (handlers *Handlers) AuthLogin(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	returnTo := req.URL.Query().Get("return_to")

	if ok := auth.ValidateLoginRedirect(returnTo); !ok {
		HandleCustomError(res, customerrors.InvalidReturnToURL)
		return
	}

	state, err := auth.NewAuthState(returnTo)
	if err != nil {
		HandleError(res, req, err)
		return
	}
	stateStr, err := utils.EncodeToBase64(state)
	if err != nil {
		HandleError(res, req, err)
		return
	}

	// Save state inside session
	sessionManager := sessionmanager.Get()
	sessionManager.Put(ctx, auth.StateSessionKey, state)

	http.Redirect(res, req, auth.GetAuthCodeUrl(stateStr), http.StatusTemporaryRedirect)
}

func (handlers *Handlers) AuthCallback(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	sessionManager := sessionmanager.Get()

	sessionState, ok := sessionManager.Get(ctx, auth.StateSessionKey).(auth.AuthState)
	if !ok {
		HandleCustomError(res, customerrors.InvalidSessionStateError)
		return
	}
	queryParams := req.URL.Query()
	queryState := &auth.AuthState{}
	if err := utils.DecodeFromBase64(queryState, queryParams.Get("state")); err != nil {
		HandleError(res, req, err)
		return
	}

	// Verify state is valid
	if sessionState.Checksum != queryState.Checksum || sessionState.ReturnTo != queryState.ReturnTo {
		HandleCustomError(res, customerrors.InvalidStateParameterError)
		return
	}

	token, err := auth.ExchangeCodeForToken(ctx, queryParams.Get("code"))
	if err != nil {
		HandleCustomError(res, customerrors.ExchangeFailedError)
		return
	}

	idToken, err := auth.VerifyIDToken(ctx, token)
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

	http.Redirect(res, req, sessionState.ReturnTo, http.StatusTemporaryRedirect)
}

func (handlers *Handlers) AuthLogout(res http.ResponseWriter, req *http.Request) {
	sessionManager := sessionmanager.Get()
	logoutUrl, err := auth.NewLogoutUrl()
	if err != nil {
		HandleError(res, req, err)
		return
	}

	// Destroy session
	sessionManager.Destroy(req.Context())

	// Get the return url
	returnTo := req.URL.Query().Get("return_to")
	if ok := auth.ValidateLogoutRedirect(returnTo); !ok {
		HandleCustomError(res, customerrors.InvalidReturnToURL)
		return
	}

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo)
	parameters.Add("client_id", auth.GetClientId())
	logoutUrl.RawQuery = parameters.Encode()

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
