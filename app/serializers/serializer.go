package serializers

import (
	"golang_gin/app/ginapp_2/model"
)

type BaseResponse struct {
	data any `json:"data"`
}

func User(user *model.Users) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Name:      &user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
