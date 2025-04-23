package main

import (
	"database/sql"
	"golang_gin/config"
	. "golang_gin/routes"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

var conn sql.DB
var PasetoSecret string

func runMigrations() {
	dsn := config.GetDsn()
	m, err := migrate.New(
		"file://db/migrations",
		"mysql://"+dsn,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

func runEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	config.InitEnv()
	runMigrations()

	conn, err := sql.Open("mysql", config.GetDsn())

	if err != nil {
		log.Fatal(err)
	}

	Route := NewRoute(os.Getenv("APP_PORT"))

	Route.SetupRoutes(conn)
}
