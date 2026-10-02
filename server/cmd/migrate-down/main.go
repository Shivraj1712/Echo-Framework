package main

import (
	"errors"
	"log/slog"
	"os"

	"github.com/Shivraj1712/Echo-Framework/config"
	"github.com/golang-migrate/migrate/v4"
)

func MigrateUp() {
	cfg, err := config.FetchConfig()
	if err != nil {
		slog.Error("failed to load environment variables", "error", err)
		os.Exit(1)
	}
	slog.Info("Migrations Starts")
	mig, err := migrate.New("file://migrations", cfg.DatabaseURL)
	if err != nil {
		slog.Error("failed to migrate", "error", err)
	}
	if err := mig.Down(); err != nil && errors.Is(err, migrate.ErrNoChange) {
		slog.Error("failed to migrate", "error", err)
		os.Exit(1)
	}
	slog.Info("Migrations Successful")
}
