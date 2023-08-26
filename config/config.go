package config

import (
	"os"
	"strconv"
)

type Config struct {
	DBHost        string
	DBUser        string
	DBPassword    string
	DBName        string
	DBPort        string
	JWTPrivateKey string
	JWTTTL        int
}

func NewConfig() *Config {
	return &Config{
		DBHost:        getEnvOrDefault("DB_HOST", "localhost"),
		DBUser:        getEnvOrDefault("DB_USER", "username"),
		DBPassword:    getEnvOrDefault("DB_PASSWORD", "password"),
		DBName:        getEnvOrDefault("DB_NAME", "dbname"),
		DBPort:        getEnvOrDefault("DB_PORT", "5432"),
		JWTPrivateKey: getEnvOrDefault("JWT_PRIVATE_KEY", "your-secret-key"),
		JWTTTL:        parseIntEnvOrDefault("JWT_TTL", 30),
	}
}

func getEnvOrDefault(envName string, defaultValue string) string {
	value := os.Getenv(envName)
	if value == "" {
		return defaultValue
	}
	return value
}

func parseIntEnvOrDefault(envName string, defaultValue int) int {
	valStr := os.Getenv(envName)
	if valStr == "" {
		return defaultValue
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultValue
	}
	return val
}
