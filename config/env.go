package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var PasetoSecret string

func InitEnv() {
	_ = godotenv.Load()
	PasetoSecret = os.Getenv("PASETO_SECRET")
	if len(PasetoSecret) != 32 {
		log.Fatal("PASETO_SECRET must be 32 characters")
	}
}
