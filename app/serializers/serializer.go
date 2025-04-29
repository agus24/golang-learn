package serializers

import (
	"golang_gin/app/databases/model"
	"golang_gin/app/repositories"
	"golang_gin/app/utils"
)

func Pagination(page, perPage *int64) PaginationResponse {
	return PaginationResponse{
		Page:    *page,
		PerPage: *perPage,
	}
}

func ValidationError(err error) ValidationErrorResponse {
	return ValidationErrorResponse{
		Message:         "Invalid Input",
		ErrorCode:       "validation_failed",
		ValidationError: utils.GenerateValidationErrors(err),
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

	if len(result) == 0 {
		return []CategoryResponse{}
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

	if len(result) == 0 {
		return []SubCategoryResponse{}
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

	if len(result) == 0 {
		return []ItemResponse{}
	}

	return result
}

func Order(order *model.Orders) OrderResponse {
	return OrderResponse{
		ID:           order.ID,
		Date:         order.Date,
		OrderNumber:  order.OrderNumber,
		GrandTotal:   order.GrandTotal,
		CustomerName: order.CustomerName,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
	}
}

func Orders(orders []model.Orders) []OrderResponse {
	var result []OrderResponse
	for _, order := range orders {
		result = append(result, Order(&order))
	}

	if len(result) == 0 {
		return []OrderResponse{}
	}

	return result
}
