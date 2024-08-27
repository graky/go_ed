package config

import (
	"os"
)

type Config struct {
	DatabaseURL   string
	ServerAddress string
	JWTSecret     string
}

func Load() *Config {
	return &Config{
		DatabaseURL:   getEnv("DATABASE_URL", "postgres://postgre:postgre@localhost/go_ed?sslmode=disable"),
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
		JWTSecret:     getEnv("JWT_SECRET", "secret_key"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}