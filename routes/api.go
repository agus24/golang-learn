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

	app := v1.Group("/app", middlewares.AuthMiddleware)
	setupAppRoute(app)
}

func setupAuthRoutes(r *gin.RouterGroup) {
	authController := controllers.NewAuthController(db)
	userController := controllers.NewUserController(db)

	auth := r.Group("/auth")

	auth.POST("/login", authController.Login)
	auth.GET("/user", middlewares.AuthMiddleware, authController.User)
	auth.POST("/user", middlewares.AuthMiddleware, userController.CreateUser)
}

func setupAppRoute(r *gin.RouterGroup) {
	categoryController := controllers.NewCategoryController(db)
	subCategoryController := controllers.NewSubCategoryController(db)

	category := r.Group("/categories")
	{
		category.GET("", categoryController.GetAllCategories)
		category.POST("", categoryController.CreateCategory)
		category.GET(":id", categoryController.GetCategory)
		category.PUT(":id", categoryController.UpdateCategory)
		category.DELETE(":id", categoryController.DeleteCategory)
	}

	subCategory := r.Group("/sub-categories")
	{
		subCategory.GET("", subCategoryController.GetAllSubCategories)
		subCategory.POST("", subCategoryController.CreateSubCategory)
		subCategory.GET(":id", subCategoryController.GetSubCategory)
		subCategory.PUT(":id", subCategoryController.UpdateSubCategory)
		subCategory.DELETE(":id", subCategoryController.DeleteSubCategory)
	}
}
