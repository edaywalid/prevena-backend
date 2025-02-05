package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ENV                 string
	PORT                string
	MONGO_URI           string
	DB_NAME             string
	PROD_URL            string
	DOCS_URL            string
	CORS_ALLOWED_ORIGIN string
	REDIS_ADDR          string
}

func (c *Config) IsProduction() bool {
	return c.ENV == "production"
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
	config.DB_NAME = getEnv("DB_NAME", "pinktober")
	config.MONGO_URI = getEnv("MONGO_URI", "mongodb://localhost:27017")
	config.PROD_URL = getEnv("PROD_URL", "")
	config.DOCS_URL = getEnv("DOCS_URL", "")
	config.CORS_ALLOWED_ORIGIN = getEnv("CORS_ALLOWED_ORIGIN", "")
	config.REDIS_ADDR = getEnv("REDIS_ADDR", "localhost:6379")
	return &config, nil
}
