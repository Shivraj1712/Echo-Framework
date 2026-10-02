package main

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Shivraj1712/Echo-Framework/config"
	"github.com/Shivraj1712/Echo-Framework/database"
	"github.com/Shivraj1712/Echo-Framework/server"
	"github.com/labstack/echo/v5"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	config, err := config.FetchConfig()
	if err != nil {
		slog.Error("failed to get the environments ready", "error", err)
		os.Exit(1)
	}
	if config.DatabaseURL == "" {
		slog.Error("no connection string for the database connection", "error", errors.New("No Database connection found in env"))
		os.Exit(1)
	}
	db, err := database.ConnectDB(context.Background(), config.DatabaseURL)
	if err != nil {
		slog.Error("failed to connect to the database", "error", err)
		os.Exit(1)
	}
	defer db.Close()
	server := server.New()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	start := echo.StartConfig{
		Address:         ":8080",
		GracefulTimeout: 20 * time.Second,
	}
	err = start.Start(ctx, server)
	if err != nil {
		slog.Error("Failed to start the server", "error", err.Error())
		os.Exit(1)
	}
}
