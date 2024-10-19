package config

import "github.com/joho/godotenv"

type Config struct {
	ENV  string
	PORT string
}

func LoadConfig() (*Config, error) {
	var config Config
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	config.ENV = getEnv("ENV", "development")
	config.PORT = getEnv("PORT", "8080")

	return &config, nil
}

func getEnv(key string, defaultValue string) string {
	value, exists := LookupEnv(key)
	if exists {
		return value
	}

	return defaultValue
}

func LookupEnv(key string) (string, bool) {
	if env_map, err := godotenv.Read(); err == nil {
		if value, exists := env_map[key]; exists {
			return value, true
		}
	}
	return "", false
}
