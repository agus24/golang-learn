package controllers

import (
	"database/sql"
	"golang_gin/app/repositories"
	"golang_gin/app/serializers"
	"golang_gin/app/services"
	"golang_gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type CategoryController struct {
	service *services.CategoryService
}

func NewCategoryController(db *sql.DB) *CategoryController {
	return &CategoryController{services.NewCategoryService(repositories.NewCategoryRepository(db))}
}

func (self *CategoryController) GetAllCategories(ctx *gin.Context) {
	search := ctx.DefaultQuery("search", "")
	pageQuery := ctx.DefaultQuery("page", "1")
	perPageQuery := ctx.DefaultQuery("per_page", "")

	page, perPage, err := utils.ParsePageAndPerPage(pageQuery, perPageQuery)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid page or per_page value"})
		return
	}

	categories, err := self.service.GetAllCategories(search, *page, *perPage)

	utils.Handle(ctx, func() []serializers.CategoryResponse {
		return serializers.Categories(categories)
	}, err, http.StatusCreated)
}

func (self *CategoryController) CreateCategory(ctx *gin.Context) {
	var input CreateCategoryRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	category, err := self.service.CreateCategory(input.Name)
	utils.Handle(ctx, func() serializers.CategoryResponse {
		return serializers.Category(category)
	}, err, http.StatusCreated)
}
