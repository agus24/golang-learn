package routes

import (
	"database/sql"
	"golang_gin/app/middlewares"

	"github.com/gin-gonic/gin"
)

type Route struct{}

func NewRoute() *Route {
	return &Route{}
}

func (route *Route) SetupRoutes(conn *sql.DB) *gin.Engine {
	r := gin.Default()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(middlewares.CorsMiddleware())
	// r.Use(middlewares.RateLimiterMiddleware())

	r.GET("/__health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})
	})

	SetupApiRoutes(r, conn)

	return r
}
