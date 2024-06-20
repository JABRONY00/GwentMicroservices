package helpers

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Panicf("Error loading .env file: %v", err)
	}
	return os.Getenv(key)
}

func CheckRequiredEnvs() {
	requiredEnvVars := []string{
		"SERVER_PORT",
		"SERVER_HOST",

		"LOG_LEVEL",

		"DB_PORT",
		"DB_HOST",
		"DB_NAME",
		"DB_USER",
		"DB_PASSWORD",
	}

	for _, envVar := range requiredEnvVars {
		if value, exists := os.LookupEnv(envVar); !exists || value == "" {
			log.Panic(fmt.Sprintf("Error: Environment variable %v is not set.", envVar))
		}
	}
}
