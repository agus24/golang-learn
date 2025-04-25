package config

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func GenerateDsn() *string {
	format := "%s:%s@tcp(%s:%s)/%s"

	dsn := fmt.Sprintf(format,
		DbUser,
		DbPass,
		DbHost,
		DbPort,
		DbName,
	)

	return &dsn
}

func RunMigrations() {
	dsn := *GenerateDsn()
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
