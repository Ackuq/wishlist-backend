package db

import (
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(connectionString string) error {
	slog.Info("Migrating DB...")

	m, err := migrate.New("file://internal/db/migrations", connectionString)

	if err != nil {
		return err
	}

	if err = m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			slog.Info("No migrations to apply")
			return nil
		}
		return err
	}

	slog.Info("Successfully applied all migrations")
	return nil
}
