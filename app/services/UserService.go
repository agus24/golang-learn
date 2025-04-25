package services

import (
	"errors"
	"golang_gin/app/databases/model"
	. "golang_gin/app/libraries"
	"golang_gin/app/repositories"
	"golang_gin/app/requests"
	"golang_gin/utils"
)

type UserService struct {
	Repo *repositories.UserRepository
}

type LoginResult struct {
	User  *model.Users
	Token *string
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (self UserService) LoginUser(username string, password string) (LoginResult, error) {
	var token string
	user, err := self.Repo.GetUserByUsername(username)
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

func (self UserService) GetUserById(id int64) (*model.Users, error) {
	return self.Repo.GetUserById(id)
}

func (self UserService) CreateUser(data requests.UserCreateRequest) (*model.Users, error) {
	password, err := utils.HashPassword(data.Password)

	if err != nil {
		return nil, err
	}

	result, err := self.Repo.CreateUser(data.Username, password, data.Name)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	return self.Repo.GetUserById(id)
}
