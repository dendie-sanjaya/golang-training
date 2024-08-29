package config

import (
	"os"
	"strconv"
)

// this is just a sample. in production, should put into kubernetes secret / .env file
var (
	KafkaBroker        = getEnv("KAFKA_BROKER", "localhost:9092")
	KafkaGroupID       = getEnv("KAFKA_GROUP_ID", "user-group-1")
	KafkaTopicTransfer = getEnv("KAFKA_TOPIC", "transaction_transfer")
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
