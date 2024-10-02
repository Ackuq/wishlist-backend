package config

import (
	"flag"
	"os"
)

type Config struct {
	Host     string
	Database struct {
		URL string
	}
}

func GetConfig() *Config {
	config := &Config{}

	flag.StringVar(&config.Host, "host", os.Getenv("HOST"), "API server host")

	flag.StringVar(&config.Database.URL, "db-url", os.Getenv("DB_URL"), "Database name")

	return config
}
