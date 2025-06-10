package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI     string
	DatabaseName string
	JWTSecret    string
	Port         string
}

func LoadConfig() *Config {
	// Загружаем .env файл, если он существует
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config := &Config{
		MongoURI:     getEnv("MONGO_URI", "mongodb://admin:secure_password@localhost:27017/trucking_db?authSource=admin"),
		DatabaseName: getEnv("DATABASE_NAME", "trucking_db"),
		JWTSecret:    getEnv("JWT_SECRET", "your_jwt_secret_here_change_in_production"),
		Port:         getEnv("PORT", "8080"),
	}

	return config
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
