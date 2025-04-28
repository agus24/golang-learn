package requests

type UserCreateRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type SubCategoryCreateRequest struct {
	Name       string `json:"name" binding:"required"`
	CategoryID int64  `json:"category_id" binding:"required"`
}

type SubCategoryUpdateRequest struct {
	Name       string `json:"name" binding:"required"`
	CategoryID int64  `json:"category_id" binding:"required"`
}

type ItemCreateRequest struct {
	Name          string `json:"name" binding:"required"`
	Price         int    `json:"price" binding:"required"`
	SubCategoryID int64  `json:"sub_category_id" binding:"required"`
}

type ItemUpdateRequest struct {
	Name          string `json:"name" binding:"required"`
	Price         int    `json:"price" binding:"required"`
	SubCategoryID int64  `json:"sub_category_id" binding:"required"`
}

type OrderRequest struct {
	CustomerName string `json:"customer_name" binding:"required"`
	Date         string `json:"date" binding:"required"`
	Details      []struct {
		ItemID   int64 `json:"item_id" binding:"required"`
		Quantity int   `json:"quantity" binding:"required"`
	}
}
