package routes

import (
	"database/sql"
	"golang_gin/app/controllers"
	"golang_gin/app/middlewares"
	"golang_gin/app/requests"

	"github.com/gin-gonic/gin"
)

func SetupApiRoutes(r *gin.Engine, conn *sql.DB) *gin.RouterGroup {
	v1 := r.Group("/api/v1")

	setupAuthRoutes(v1, conn)

	app := v1.Group("/app", middlewares.AuthMiddleware)
	setupAppRoute(app, conn)

	return v1
}

func setupAuthRoutes(r *gin.RouterGroup, db *sql.DB) {
	authController := controllers.NewAuthController(db)
	userController := controllers.NewUserController(db)

	auth := r.Group("/auth")

	auth.POST("/login", requests.BasicValidation[requests.LoginRequest], authController.Login)
	auth.GET("/user", middlewares.AuthMiddleware, authController.User)
	auth.POST("/user", middlewares.AuthMiddleware, userController.CreateUser)
}

func setupAppRoute(r *gin.RouterGroup, db *sql.DB) {
	categoryController := controllers.NewCategoryController(db)
	subCategoryController := controllers.NewSubCategoryController(db)
	itemController := controllers.NewItemController(db)
	orderController := controllers.NewOrderController(db)

	category := r.Group("/categories")
	{
		category.GET("", categoryController.GetAllCategories)
		category.POST("", requests.BasicValidation[requests.CreateCategoryRequest], categoryController.CreateCategory)
		category.GET(":id", categoryController.GetCategory)
		category.PUT(":id", requests.BasicValidation[requests.UpdateCategoryRequest], categoryController.UpdateCategory)
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

	item := r.Group("/items")
	{
		item.GET("", itemController.GetAll)
		item.POST("", itemController.Create)
		item.GET(":id", itemController.Show)
		item.PUT(":id", itemController.Update)
		item.DELETE(":id", itemController.Delete)
	}

	order := r.Group("/orders")
	{
		order.GET("", orderController.GetAll)
		order.POST("", requests.BasicValidation[requests.CreateOrderRequest], orderController.Create)
		order.GET(":id", orderController.Show)
	}
}
