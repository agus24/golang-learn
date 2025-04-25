package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var PasetoSecret string
var PasetoExpirationTime time.Duration

var DbUser string
var DbPass string
var DbHost string
var DbPort string
var DbName string

func InitEnv(envFile string) {
	err := godotenv.Load(envFile)

	if err != nil {
		log.Fatal("Error loading .env file: " + envFile + " " + err.Error())
	}

	parseDatabase()
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

func parseDatabase() {
	DbUser = os.Getenv("DB_USER")
	DbPass = os.Getenv("DB_PASS")
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbName = os.Getenv("DB_NAME")
}
