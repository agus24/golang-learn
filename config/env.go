package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var PasetoSecret string
var PasetoExpirationTime time.Duration

func InitEnv() {
	_ = godotenv.Load()

	parsePasetoSecret()
	parsePasetoExpirationTime()
}

func parsePasetoSecret() {
	PasetoSecret = os.Getenv("PASETO_SECRET")
	if len(PasetoSecret) != 32 {
		log.Fatal("PASETO_SECRET must be 32 characters")
	}
}

func parsePasetoExpirationTime() {
	var err error
	expirationTime := os.Getenv("PASETO_EXPIRATION_TIME")

	PasetoExpirationTime, err = time.ParseDuration(expirationTime + "m")

	if err != nil {
		log.Fatal("PASETO_EXPIRATION_TIME must be a valid duration")
	}
}
