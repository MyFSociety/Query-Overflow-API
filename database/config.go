package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not load environment variables")
	}

	return os.Getenv("MONGO_URL")
}
