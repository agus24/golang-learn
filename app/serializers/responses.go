package serializers

import "time"

type UserResponse struct {
	ID        int64      `json:"id"`
	Username  string     `json:"username"`
	Name      *string    `json:"name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
