package controllers

import (
	"database/sql"
	"golang_gin/app/repositories"
	"golang_gin/app/serializers"
	"golang_gin/app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *services.UserService
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewAuthController(db *sql.DB) *AuthController {
	service := services.NewUserService(repositories.NewUserRepository(db))
	return &AuthController{
		service,
	}
}

func (self AuthController) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req LoginRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		output, err := self.service.LoginUser(ctx, req.Username, req.Password)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"user": output.User, "token": output.Token})
	}
}

func (self AuthController) User() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		val := ctx.MustGet("userID").(string)

		userID, err := strconv.Atoi(val)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
			return
		}

		user, err := self.service.GetUserById(ctx, int64(userID))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": serializers.User(user)})
	}
}
