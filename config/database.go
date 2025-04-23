package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var dsn *string

func GetDsn() string {
	if dsn == nil {
		dsn = generateDsn()
	}

	return *dsn
}

func generateDsn() *string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	format := "%s:%s@tcp(%s:%s)/%s"

	dsn := fmt.Sprintf(format,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	return &dsn
}
