package serializers

import (
	"golang_gin/app/ginapp_2/model"
)

func Pagination(page, perPage *int64) PaginationResponse {
	return PaginationResponse{
		Page:    *page,
		PerPage: *perPage,
	}
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

func Category(category *model.Categories) CategoryResponse {
	return CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}

func Categories(categories []model.Categories) []CategoryResponse {
	var result []CategoryResponse
	for _, category := range categories {
		result = append(result, Category(&category))
	}

	return result
}
