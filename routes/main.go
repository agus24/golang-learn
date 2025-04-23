package routes

import (
	"database/sql"
	"golang_gin/app/middlewares"

	"github.com/gin-gonic/gin"
)

type Route struct {
	port string
}

func NewRoute(port string) *Route {
	return &Route{port}
}

func (route *Route) SetupRoutes(conn *sql.DB) {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(middlewares.CorsMiddleware())
	r.Use(middlewares.RateLimiterMiddleware())

	SetupApiRoutes(r, conn)

	r.Run(":" + route.port)
}
