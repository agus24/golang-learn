package controllers

import (
	"database/sql"
	"golang_gin/app/repositories"
	"golang_gin/app/requests"
	"golang_gin/app/serializers"
	"golang_gin/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(db *sql.DB) *UserController {
	return &UserController{
		services.NewUserService(repositories.NewUserRepository(db)),
	}
}

func (self UserController) CreateUser(ctx *gin.Context) {
	var req requests.UserCreateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := self.service.CreateUser(ctx, req)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed create user"})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": serializers.User(user)})
}
