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

type SubCategoryCreateRequest struct {
	Name       string `json:"name" binding:"required"`
	CategoryID int64  `json:"category_id" binding:"required"`
}

type SubCategoryUpdateRequest struct {
	Name       string `json:"name" binding:"required"`
	CategoryID int64  `json:"category_id" binding:"required"`
}

type SubCategoryController struct {
	service         *services.SubCategoryService
	categoryService *services.CategoryService
}

func NewSubCategoryController(db *sql.DB) *SubCategoryController {
	return &SubCategoryController{
		service:         services.NewSubCategoryService(repositories.NewSubCategoryRepository(db)),
		categoryService: services.NewCategoryService(repositories.NewCategoryRepository(db)),
	}
}

func (self *SubCategoryController) getByParam(ctx *gin.Context, id string) (*repositories.SubCategory, error) {
	subCategoryID := utils.ParseToInt(ctx, ctx.Param("id"))

	subCategory, err := self.service.GetSubCategoryById(*subCategoryID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, serializers.ErrorResponse("Sub Category not found"))
		return nil, err
	}

	return subCategory, nil
}

func (self *SubCategoryController) checkCategoryExists(ctx *gin.Context, categoryID int64) bool {
	_, err := self.categoryService.Repo.GetCategoryById(categoryID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, serializers.ErrorResponse("Category not found"))
		return false
	}

	return true
}

func (self *SubCategoryController) GetAllSubCategories(ctx *gin.Context) {
	search := ctx.DefaultQuery("search", "")
	pageQuery := ctx.DefaultQuery("page", "1")
	perPageQuery := ctx.DefaultQuery("per_page", "")

	page, perPage, err := utils.ParsePageAndPerPage(pageQuery, perPageQuery)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid page or per_page value"})
		return
	}

	subCategories, err := self.service.GetAllSubCategories(search, *page, *perPage)

	utils.Handle(ctx, func() gin.H {
		return gin.H{"data": serializers.SubCategories(subCategories), "meta": serializers.Pagination(page, perPage)}
	}, err, http.StatusOK)
}

func (self *SubCategoryController) GetSubCategory(ctx *gin.Context) {
	subCategory, err := self.getByParam(ctx, ctx.Param("id"))

	utils.Handle(ctx, func() gin.H {
		return gin.H{"data": serializers.SubCategory(subCategory)}
	}, err, http.StatusOK)
}

func (self *SubCategoryController) CreateSubCategory(ctx *gin.Context) {
	var input SubCategoryCreateRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, serializers.ErrorResponse(err.Error()))
		return
	}

	if !self.checkCategoryExists(ctx, input.CategoryID) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, serializers.ErrorResponse("Category not found"))
	}

	subCategory, err := self.service.CreateSubCategory(input.Name, input.CategoryID)

	utils.Handle(ctx, func() gin.H {
		return gin.H{"data": serializers.SubCategory(subCategory)}
	}, err, http.StatusOK)
}

func (self *SubCategoryController) UpdateSubCategory(ctx *gin.Context) {
	var input SubCategoryUpdateRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, serializers.ErrorResponse(err.Error()))
		return
	}

	self.checkCategoryExists(ctx, input.CategoryID)

	subCategory, err := self.getByParam(ctx, ctx.Param("id"))
	subCategory, err = self.service.UpdateSubCategory(subCategory.ID, input.Name, input.CategoryID)

	utils.Handle(ctx, func() gin.H {
		return gin.H{"data": serializers.SubCategory(subCategory)}
	}, err, http.StatusOK)
}

func (self *SubCategoryController) DeleteSubCategory(ctx *gin.Context) {
	subCategory, err := self.getByParam(ctx, ctx.Param("id"))

	err = self.service.DeleteSubCategory(subCategory.ID)

	utils.Handle(ctx, func() gin.H {
		return gin.H{"data": "Sub Category deleted successfully"}
	}, err, http.StatusOK)
}
