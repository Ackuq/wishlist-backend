package auth

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/ackuq/wishlist-backend/internal/config"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

var (
	provider             *oidc.Provider
	authConfig           oauth2.Config
	logoutUrl            string
	validLoginRedirects  []string
	validLogoutRedirects []string
)

type Claims struct {
	// OpenID scope
	Issuer   string  `json:"iss"`
	Subject  string  `json:"sub"`
	Audience string  `json:"aud"`
	Expiry   float64 `json:"exp"`
	IssuedAt float64 `json:"iat"`
	// Profile
	Name       string `json:"name"`
	FamilyName string `json:"family_name"`
	GivenName  string `json:"given_name"`
	Nickname   string `json:"nickname"`
	Picture    string `json:"picture"`
	// Email scope
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

const StateSessionKey = "state"
const AccessTokenSessionKey = "access_token"
const ClaimsSessionKey = "claims"

func Init(config *config.Config) error {
	_provider, err := oidc.NewProvider(
		context.Background(),
		fmt.Sprintf("https://%s/", config.Auth0.Domain),
	)
	if err != nil {
		return err
	}
	provider = _provider

	authConfig = oauth2.Config{
		ClientID:     config.Auth0.ClientID,
		ClientSecret: config.Auth0.ClientSecret,
		RedirectURL:  config.Auth0.CallbackURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	logoutUrl = fmt.Sprintf("https://%s/v2/logout", config.Auth0.Domain)
	validLoginRedirects = config.Redirects.ValidLoginRedirects
	validLogoutRedirects = config.Redirects.ValidLogoutRedirects

	return nil
}

func VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: authConfig.ClientID,
	}

	return provider.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}

func ExchangeCodeForToken(ctx context.Context, code string) (*oauth2.Token, error) {
	return authConfig.Exchange(ctx, code)
}

func GetClientId() string {
	return authConfig.ClientID
}

func NewLogoutUrl() (*url.URL, error) {
	return url.Parse(fmt.Sprintf("https://%s/v2/logout", logoutUrl))
}

func GetAuthCodeUrl(state string) string {
	return authConfig.AuthCodeURL(state)
}
