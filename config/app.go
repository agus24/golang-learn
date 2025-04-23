package config

import (
	"os"

	"github.com/gin-gonic/gin"
)

var DefaultPerPage int64
var Debug bool = false

func InitConfig() {
	DefaultPerPage = 15
	Debug = os.Getenv("APP_DEBUG") == "true"
}
