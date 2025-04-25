package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"golang_gin/config"
	"golang_gin/routes"
	"golang_gin/utils"
	"io"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

var Engine *gin.Engine

func CreateApplication() {
	config.InitEnv("../.env.test")
	config.InitConfig()
	config.RunMigrations()

	conn, err := sql.Open("mysql", config.GetDsn())

	utils.Dump(conn, true)

	if err != nil {
		log.Fatal(err)
	}

	route := routes.NewRoute()

	Engine = route.SetupRoutes(conn)
}

func GetTestServer() *httptest.Server {
	if Engine == nil {
		CreateApplication()
	}

	return httptest.NewServer(Engine)
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
