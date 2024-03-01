package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("error .env")
	}

	mongoUri := os.Getenv("MONGOURI")
	return mongoUri
}
