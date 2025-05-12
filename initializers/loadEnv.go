package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error load .env file" + err.Error())
	}
}
