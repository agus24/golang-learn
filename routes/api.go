package routes

import (
	"database/sql"
	"golang_gin/app/controllers"
	"golang_gin/app/middlewares"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func SetupApiRoutes(r *gin.Engine, conn *sql.DB) {
	db = conn
	helloController := controllers.NewHelloController()

	v1 := r.Group("/api/v1")

	v1.GET("/hello", helloController.Hello)
	setupAuthRoutes(v1)
}

func setupAuthRoutes(r *gin.RouterGroup) {
	authController := controllers.NewAuthController(db)
	userController := controllers.NewUserController(db)

	auth := r.Group("/auth")

	auth.POST("/login", authController.Login)
	auth.GET("/user", middlewares.AuthMiddleware(), authController.User)
	auth.POST("/user", middlewares.AuthMiddleware(), userController.CreateUser)
}
