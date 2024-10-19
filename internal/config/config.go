package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ENV  string
	PORT string
}

func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func LoadConfig() (*Config, error) {
	var config Config

	if getEnv("ENV", "development") == "development" {
		if err := godotenv.Load(); err != nil {
			return nil, fmt.Errorf("Error loading .env file: %v", err)
		}
	}

	config.ENV = getEnv("ENV", "development")
	config.PORT = getEnv("PORT", "8080")

	return &config, nil
}
