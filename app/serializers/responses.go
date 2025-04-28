package serializers

import "time"

type PaginationResponse struct {
	Page    int64 `json:"page"`
	PerPage int64 `json:"per_page"`
}

type ValidationErrorResponse struct {
	Message         string            `json:"message"`
	ErrorCode       string            `json:"error_code"`
	ValidationError map[string]string `json:"validation_errors"`
}

type UserResponse struct {
	ID        int64      `json:"id"`
	Username  string     `json:"username"`
	Name      *string    `json:"name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type CategoryResponse struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type SubCategoryResponse struct {
	ID         int64             `json:"id"`
	Name       string            `json:"name"`
	CategoryID int64             `json:"category_id"`
	CreatedAt  *time.Time        `json:"created_at,omitempty"`
	UpdatedAt  *time.Time        `json:"updated_at,omitempty"`
	Category   *CategoryResponse `json:"category"`
}

type ItemResponse struct {
	ID            int64                `json:"id"`
	Name          string               `json:"name"`
	Price         int32                `json:"price"`
	SubCategoryID int64                `json:"sub_category_id"`
	CreatedAt     *time.Time           `json:"created_at,omitempty"`
	UpdatedAt     *time.Time           `json:"updated_at,omitempty"`
	SubCategory   *SubCategoryResponse `json:"sub_category"`
}

type OrderResponse struct {
	ID           int64      `json:"id"`
	Date         time.Time  `json:"date"`
	OrderNumber  string     `json:"order_number"`
	GrandTotal   int32      `json:"grand_total"`
	CustomerName string     `json:"customer_name"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}
