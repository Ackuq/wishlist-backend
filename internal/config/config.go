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
	ValidLoginRedirects []string
	Auth0               struct {
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
	config.ValidLoginRedirects = strings.Split(validLoginRedirects, ",")

	return config
}
