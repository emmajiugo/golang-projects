package env

import (
	"log"
	"os"
	"strconv"
)

func GetString(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func GetInt(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Error converting environment variable %s to int: %v. Using default value: %d", key, err, defaultValue)
		return defaultValue
	}
	return intValue
}