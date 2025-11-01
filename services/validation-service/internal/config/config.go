package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// Config holds all configuration for the validation service
type Config struct {
	Port        int    `json:"port"`
	Environment string `json:"environment"`
	LogLevel    string `json:"log_level"`

	// Database configuration
	DatabaseURL      string `json:"database_url"`
	DatabaseHost     string `json:"database_host"`
	DatabasePort     int    `json:"database_port"`
	DatabaseName     string `json:"database_name"`
	DatabaseUser     string `json:"database_user"`
	DatabasePassword string `json:"database_password"`

	// Redis configuration
	RedisURL      string `json:"redis_url"`
	RedisHost     string `json:"redis_host"`
	RedisPort     int    `json:"redis_port"`
	RedisPassword string `json:"redis_password"`
	RedisDB       int    `json:"redis_db"`

	// Service configuration
	ServiceName    string `json:"service_name"`
	ServiceVersion string `json:"service_version"`
}

// Load loads configuration from environment variables
func Load() *Config {
	// Load .env file if it exists (for development)
	_ = godotenv.Load()

	cfg := &Config{
		Port:        getEnvAsInt("PORT", 8081),
		Environment: getEnv("ENVIRONMENT", "development"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),

		// Database
		DatabaseURL:      getEnv("DATABASE_URL", ""),
		DatabaseHost:     getEnv("DB_HOST", "localhost"),
		DatabasePort:     getEnvAsInt("DB_PORT", 5432),
		DatabaseName:     getEnv("DB_NAME", "gtrs_validation"),
		DatabaseUser:     getEnv("DB_USER", "gtrs_user"),
		DatabasePassword: getEnv("DB_PASSWORD", "gtrs_password"),

		// Redis
		RedisURL:      getEnv("REDIS_URL", ""),
		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnvAsInt("REDIS_PORT", 6379),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvAsInt("REDIS_DB", 0),

		// Service
		ServiceName:    getEnv("SERVICE_NAME", "validation-service"),
		ServiceVersion: getEnv("SERVICE_VERSION", "1.0.0-SNAPSHOT"),
	}

	// Build database URL if not provided
	if cfg.DatabaseURL == "" {
		cfg.DatabaseURL = buildDatabaseURL(cfg)
	}

	// Build Redis URL if not provided
	if cfg.RedisURL == "" {
		cfg.RedisURL = buildRedisURL(cfg)
	}

	logrus.WithFields(logrus.Fields{
		"port":        cfg.Port,
		"environment": cfg.Environment,
		"log_level":   cfg.LogLevel,
		"service":     cfg.ServiceName,
		"version":     cfg.ServiceVersion,
	}).Info("Configuration loaded")

	return cfg
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func buildDatabaseURL(cfg *Config) string {
	return "postgres://" + cfg.DatabaseUser + ":" + cfg.DatabasePassword +
		"@" + cfg.DatabaseHost + ":" + strconv.Itoa(cfg.DatabasePort) +
		"/" + cfg.DatabaseName + "?sslmode=disable"
}

func buildRedisURL(cfg *Config) string {
	if cfg.RedisPassword != "" {
		return "redis://:" + cfg.RedisPassword + "@" + cfg.RedisHost + ":" + strconv.Itoa(cfg.RedisPort) + "/" + strconv.Itoa(cfg.RedisDB)
	}
	return "redis://" + cfg.RedisHost + ":" + strconv.Itoa(cfg.RedisPort) + "/" + strconv.Itoa(cfg.RedisDB)
}