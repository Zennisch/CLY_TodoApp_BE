package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host               string
	Port               string
	CorsAllowedOrigins string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	return &Config{
		Host:               getEnv("HOST", "localhost"),
		Port:               getEnv("PORT", "8000"),
		CorsAllowedOrigins: getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:3000"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
