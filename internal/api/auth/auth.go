package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/ackuq/wishlist-backend/internal/config"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	*oidc.Provider
	oauth2.Config
	LogoutUrl string
}

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

func New(config *config.Config) (*Authenticator, error) {

	provider, err := oidc.NewProvider(
		context.Background(),
		fmt.Sprintf("https://%s/", config.Auth0.Domain),
	)
	if err != nil {
		return nil, err
	}

	authConfig := oauth2.Config{
		ClientID:     config.Auth0.ClientID,
		ClientSecret: config.Auth0.ClientSecret,
		RedirectURL:  config.Auth0.CallbackURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	logoutUrl := fmt.Sprintf("https://%s/v2/logout", config.Auth0.Domain)

	return &Authenticator{provider, authConfig, logoutUrl}, nil
}

func (auth *Authenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: auth.ClientID,
	}

	return auth.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}
