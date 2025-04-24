package serializers

import (
	"golang_gin/app/ginapp_2/model"
	"golang_gin/app/repositories"
)

func Pagination(page, perPage *int64) PaginationResponse {
	return PaginationResponse{
		Page:    *page,
		PerPage: *perPage,
	}
}

func ErrorResponse(message string) map[string]string {
	return map[string]string{
		"message": message,
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

func SubCategory(subCategory *repositories.SubCategory) SubCategoryResponse {
	var categoryResponse *CategoryResponse
	if subCategory.Category != nil {
		res := Category(subCategory.Category)
		categoryResponse = &res
	}

	return SubCategoryResponse{
		ID:         subCategory.ID,
		Name:       subCategory.Name,
		CategoryID: subCategory.CategoryID,
		CreatedAt:  subCategory.CreatedAt,
		UpdatedAt:  subCategory.UpdatedAt,
		Category:   categoryResponse,
	}
}

func SubCategories(subCategories []repositories.SubCategory) []SubCategoryResponse {
	var result []SubCategoryResponse
	for _, subCategory := range subCategories {
		result = append(result, SubCategory(&subCategory))
	}

	return result
}

func Item(item *repositories.Item) ItemResponse {
	var subCategoryResponse *SubCategoryResponse
	if item.SubCategory != nil {
		res := SubCategory(item.SubCategory)
		subCategoryResponse = &res
	}

	return ItemResponse{
		ID:            item.ID,
		Name:          item.Name,
		Price:         item.Price,
		SubCategoryID: item.SubCategoryID,
		CreatedAt:     item.CreatedAt,
		UpdatedAt:     item.UpdatedAt,
		SubCategory:   subCategoryResponse,
	}
}

func Items(items []repositories.Item) []ItemResponse {
	var result []ItemResponse
	for _, item := range items {
		result = append(result, Item(&item))
	}

	return result
}
