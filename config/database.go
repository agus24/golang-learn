package config

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var dsn *string

func GetDsn() string {
	if dsn == nil {
		dsn = generateDsn()
	}

	return *dsn
}

func generateDsn() *string {
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

func RunMigrations() {
	dsn := GetDsn()
	m, err := migrate.New(
		"file://"+os.Getenv("DB_MIGRATION_PATH"),
		"mysql://"+dsn,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
