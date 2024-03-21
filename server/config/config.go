package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	env := os.Getenv("APP_ENV")
	fmt.Println(env)
	if env == "DEVELOPMENT" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	return os.Getenv("MONGOURI")
}
