package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	AppEnv      string
	Port        string
	DatabaseURL string
	RedisURL    string
}

func FetchConfig() (*Configuration, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	appEnv := getEnv("APP_ENV", "development")
	port := getEnv("PORT", "8080")
	databaseURL := getEnv("DATABASE_URL", "")
	redisURL := getEnv("REDIS_URL", "")
	config := &Configuration{
		AppEnv:      appEnv,
		Port:        port,
		DatabaseURL: databaseURL,
		RedisURL:    redisURL,
	}
	return config, nil
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	} else {
		return value
	}
}
