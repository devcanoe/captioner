package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var env string = os.Getenv("APP_ENV")

func EnvMongoURI() string {
	if env == "DEVELOPMENT" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	return os.Getenv("MONGOURI")
}

func EnvPrivateKey() string {
	if env == "DEVELOPMENT" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	return os.Getenv("JWT_PRIVATE_KEY")
}

func EnvPublicKey() string {
	if env == "DEVELOPMENT" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	return os.Getenv("JWT_PUBLIC_KEY")
}
