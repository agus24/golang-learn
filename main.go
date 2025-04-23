package main

import (
	"database/sql"
	"golang_gin/config"
	. "golang_gin/routes"
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var conn sql.DB

func main() {
	config.InitEnv()
	config.RunMigrations()

	conn, err := sql.Open("mysql", config.GetDsn())

	if err != nil {
		log.Fatal(err)
	}

	Route := NewRoute(os.Getenv("APP_PORT"))

	Route.SetupRoutes(conn)
}
