package config

import (
	"os"
	"strconv"
)

// this is just a sample. in production, should put into kubernetes secret / .env file
var (
	AuthBasicUsername        = getEnv("AUTH_BASIC_USERNAME", "user")
	AuthBasicPassword        = getEnv("AUTH_BASIC_PASSWORD", "pass")
	RedisHost                = getEnv("REDIS_HOST", "127.0.0.1:6378")
	RedisUsername            = getEnv("REDIS_USERNAME", "")
	RedisPassword            = getEnv("REDIS_PASSWORD", "redispass")
	RedisDatabase            = getEnvAsInt("REDIS_DATABASE", 0)
	PostgresStringConnection = getEnv("POSTGRES_STRING_CONNECTION", "postgresql://postgres:password@localhost:5453/notification?sslmode=disable")
	PostgresDB               = getEnv("POSTGRES_DB", "notification")
	PostgresUser             = getEnv("POSTGRES_USER", "postgres")
	PostgresPassword         = getEnv("POSTGRES_PASSWORD", "password")
	PostgressPort            = getEnv("POSTGRES_PORT", "5452")
	PostgresSSLMode          = getEnv("POSTGRES_SSLMODE", "disable")
	PostgresHost             = getEnv("POSTGRES_HOST", "127.0.0.1")
)

// getEnv reads an environment variable and returns its value or a default value if not set.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt reads an environment variable as an integer and returns its value or a default value if not set.
func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
