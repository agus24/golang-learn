package main

import (
	"database/sql"
	"golang_gin/config"
	. "golang_gin/routes"
	"log"
	"os"
)

var conn sql.DB
var PasetoSecret string

func main() {
	config.InitEnv(".env")
	config.InitConfig()

	conn, err := sql.Open("mysql", config.GetDsn())

	if err != nil {
		log.Fatal(err)
	}

	config.RunMigrations()

	Route := NewRoute(os.Getenv("APP_PORT"))

	Route.SetupRoutes(conn)
}
