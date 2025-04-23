package controllers

import (
	"database/sql"
	"golang_gin/app/repositories"
	"golang_gin/app/services"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(db *sql.DB) *UserController {
	userService := services.NewUserService(repositories.NewUserRepository(db))
	return &UserController{
		UserService: userService,
	}
}
