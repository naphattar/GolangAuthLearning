package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURL() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}
	return os.Getenv("DatabaseURL")
}
