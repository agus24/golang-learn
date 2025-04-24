package serializers

import "time"

type PaginationResponse struct {
	Page    int64 `json:"page"`
	PerPage int64 `json:"per_page"`
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
