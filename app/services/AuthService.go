package services

import (
	"errors"
	"golang_gin/app/ginapp_2/model"
	. "golang_gin/app/libraries"
	"golang_gin/app/repositories"
	"golang_gin/utils"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	repo *repositories.UserRepository
}

type LoginResult struct {
	User  *model.Users
	Token *string
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository,
	}
}

func (self UserService) LoginUser(ctx *gin.Context, username string, password string) (LoginResult, error) {
	var token string
	user, err := self.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return LoginResult{nil, nil}, err
	}

	if !utils.HashCheck(password, user.Password) {
		return LoginResult{nil, nil}, errors.New("Invalid password")
	}

	token, err = NewPasetoToken().GenerateToken(user.ID)

	if err != nil {
		return LoginResult{nil, nil}, err
	}

	return LoginResult{user, &token}, nil
}

func (self UserService) GetUserById(ctx *gin.Context, id int64) (*model.Users, error) {
	user, err := self.repo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
