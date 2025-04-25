package main

import (
	"database/sql"
	"golang_gin/config"
	. "golang_gin/routes"
	"log"
	"os"
)

func main() {
	config.InitEnv(".env")
	config.InitConfig()

	conn, err := sql.Open("mysql", *config.GenerateDsn())

	if err != nil {
		log.Fatal(err)
	}

	config.RunMigrations()

	route := NewRoute()

	router := *route.SetupRoutes(conn)

	routerErr := router.Run(":" + os.Getenv("APP_PORT"))

	if routerErr != nil {
		log.Fatal(routerErr)
	}
}
