package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"golang_gin/config"
	"golang_gin/routes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

var Engine *gin.Engine
var Conn *sql.DB

func CreateApplication() {
	config.InitEnv("../.env.test")
	config.InitConfig()
	config.RunMigrations()

	var err error

	Conn, err = sql.Open("mysql", *config.GenerateDsn())

	if err != nil {
		log.Fatal(err)
	}

	route := routes.NewRoute()

	Engine = route.SetupRoutes(Conn)
}

func GetTestServer() (*httptest.Server, *sql.DB) {
	if Engine == nil {
		CreateApplication()
	}

	return httptest.NewServer(Engine), Conn
}

func ParseRequestBody(resp *http.Response) (map[string]any, error) {
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]any
	_ = json.Unmarshal(bodyBytes, &data)

	return data, err
}

func PrepareBody(body map[string]any) (*bytes.Buffer, error) {
	jsonBytes, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(jsonBytes), nil
}

func ResetDB(db *sql.DB) {
	var dbName string
	err := db.QueryRow("SELECT DATABASE()").Scan(&dbName)
	if err != nil {
		log.Fatalf("❌ Failed to get current DB: %v", err)
	}

	rows, err := db.Query(`
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = ?
		AND table_type = 'BASE TABLE'
	`, dbName)

	if err != nil {
		log.Fatalf("❌ Failed to fetch tables: %v", err)
	}

	defer rows.Close()

	// Temporarily disable foreign key checks
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	var table string
	for rows.Next() {
		err := rows.Scan(&table)
		if table == "schema_migrations" {
			continue
		}

		if err != nil {
			log.Fatal("❌ Failed to scan table name:", err)
		}

		_, err = db.Exec("TRUNCATE TABLE `" + table + "`")
		if err != nil {
			log.Fatalf("❌ Failed to truncate table %s: %v", table, err)
		}
		fmt.Println("🧼 Truncated:", table)
	}

	// Re-enable foreign key checks
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}
