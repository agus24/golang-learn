package controllers

import (
	"database/sql"
	"golang_gin/app/repositories"
	"golang_gin/app/requests"
	"golang_gin/app/serializers"
	"golang_gin/app/services"
	"golang_gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	utils.Handle(ctx, func() gin.H {
		return gin.H{"data": serializers.Categories(categories), "meta": serializers.Pagination(page, perPage)}
	}, err, http.StatusOK)
}

func (self *CategoryController) CreateCategory(ctx *gin.Context) {
	raw, _ := ctx.Get("validated")
	input := raw.(requests.CreateCategoryRequest)

	category, err := self.service.CreateCategory(input.Name)
	utils.Handle(ctx, func() gin.H {
		return gin.H{"data": serializers.Category(category)}
	}, err, http.StatusOK)
}

func (self *CategoryController) GetCategory(ctx *gin.Context) {
	id := utils.ParseToInt(ctx, ctx.Param("id"))

	category, err := self.service.GetCategory(*id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": serializers.Category(category)})
}

func (self *CategoryController) UpdateCategory(ctx *gin.Context) {
	id := utils.ParseToInt(ctx, ctx.Param("id"))

	raw, _ := ctx.Get("validated")
	input := raw.(requests.UpdateCategoryRequest)

	_, err := self.service.Repo.GetCategoryByName(input.Name, *id)

	if err == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Category name already exists"})
		return
	}

	category, err := self.service.UpdateCategory(*id, input.Name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": serializers.Category(category)})
}

func (self *CategoryController) DeleteCategory(ctx *gin.Context) {
	id := utils.ParseToInt(ctx, ctx.Param("id"))

	category, err := self.service.Repo.GetCategoryById(*id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = self.service.DeleteCategory(category)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
