package config

import (
	"flag"
	"os"
	"strings"
)

type Config struct {
	Host     string
	Database struct {
		URL string
	}
	CORS struct {
		AllowedOrigins []string
	}
	Redirects struct {
		ValidLoginRedirects  []string
		ValidLogoutRedirects []string
	}
	Auth0 struct {
		Domain       string
		ClientID     string
		ClientSecret string
		CallbackURL  string
	}
}

func GetConfig() *Config {
	config := &Config{}

	flag.StringVar(&config.Host, "host", os.Getenv("HOST"), "API server host")

	flag.StringVar(&config.Database.URL, "db-url", os.Getenv("DB_URL"), "Database name")

	flag.StringVar(&config.Auth0.Domain, "auth0-domain", os.Getenv("AUTH0_DOMAIN"), "Auth0 tenant domain")
	flag.StringVar(&config.Auth0.ClientID, "auth0-client-id", os.Getenv("AUTH0_CLIENT_ID"), "Auth0 client ID")
	flag.StringVar(&config.Auth0.ClientSecret, "auth0-client-secret", os.Getenv("AUTH0_CLIENT_SECRET"), "Auth0 client secret")
	flag.StringVar(&config.Auth0.CallbackURL, "auth0-callback-url", os.Getenv("AUTH0_CALLBACK_URL"), "Auth0 callback url")

	var validLoginRedirects string
	flag.StringVar(&validLoginRedirects, "valid-login-redirects", os.Getenv("VALID_LOGIN_REDIRECTS"), "Comma separated list with valid redirect locations after authentication")
	config.Redirects.ValidLoginRedirects = strings.Split(validLoginRedirects, ",")

	var validLogOutRedirects string
	flag.StringVar(&validLoginRedirects, "valid-logout-redirects", os.Getenv("VALID_LOGOUT_REDIRECTS"), "Comma separated list with valid redirect locations after logging out")
	config.Redirects.ValidLogoutRedirects = strings.Split(validLogOutRedirects, ",")

	var corsOrigins string
	flag.StringVar(&corsOrigins, "cors-origins", os.Getenv("CORS_ALLOWED_ORIGINS"), "Comma separated list of CORS origins")
	config.CORS.AllowedOrigins = strings.Split(corsOrigins, ",")

	return config
}
