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

type ItemController struct {
	service            *services.ItemService
	subCategoryService *services.SubCategoryService
}

func NewItemController(db *sql.DB) *ItemController {
	return &ItemController{
		service:            services.NewItemService(repositories.NewItemRepository(db)),
		subCategoryService: services.NewSubCategoryService(repositories.NewSubCategoryRepository(db)),
	}
}

func (self *ItemController) getByparam(ctx *gin.Context, idParam string) (*repositories.Item, error) {
	id := utils.ParseToInt(ctx, idParam)

	item, err := self.service.GetById(*id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, serializers.ErrorResponse(err.Error()))
		return nil, err
	}

	return item, nil
}

func (self *ItemController) checkSubCategoryExists(ctx *gin.Context, id int64) bool {
	_, err := self.subCategoryService.Repo.GetSubCategoryById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, serializers.ErrorResponse("Sub Category not found"))
		return false
	}

	return true
}

func (self *ItemController) GetAll(ctx *gin.Context) {
	search := ctx.DefaultQuery("search", "")
	pageQuery := ctx.DefaultQuery("page", "1")
	perPageQuery := ctx.DefaultQuery("per_page", "")

	page, perPage, err := utils.ParsePageAndPerPage(pageQuery, perPageQuery)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid page or per_page value"})
		return
	}

	items, err := self.service.GetAll(search, *page, *perPage)

	utils.Handle(ctx, func() gin.H {
		return gin.H{"data": serializers.Items(items), "meta": serializers.Pagination(page, perPage)}
	}, err, http.StatusOK)
}

func (self *ItemController) Show(ctx *gin.Context) {
	item, _ := self.getByparam(ctx, ctx.Param("id"))

	ctx.JSON(http.StatusOK, gin.H{"data": serializers.Item(item)})
}

func (self *ItemController) Create(ctx *gin.Context) {
	var request requests.ItemCreateRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, serializers.ErrorResponse(err.Error()))
		return
	}

	self.checkSubCategoryExists(ctx, request.SubCategoryID)

	item, err := self.service.Create(request)

	utils.Handle(ctx, func() gin.H {
		return gin.H{"data": serializers.Item(item)}
	}, err, http.StatusOK)
}

func (self *ItemController) Update(ctx *gin.Context) {
	var request requests.ItemUpdateRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, serializers.ErrorResponse(err.Error()))
		return
	}

	if !self.checkSubCategoryExists(ctx, request.SubCategoryID) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, serializers.ErrorResponse("Category not found"))
		return
	}

	item, err := self.getByparam(ctx, ctx.Param("id"))
	item, err = self.service.Update(item.ID, request)

	utils.Handle(ctx, func() gin.H {
		return gin.H{"data": serializers.Item(item)}
	}, err, http.StatusOK)
}

func (self *ItemController) Delete(ctx *gin.Context) {
	item, _ := self.getByparam(ctx, ctx.Param("id"))

	err := self.service.Delete(item.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializers.ErrorResponse(err.Error()))
		return
	}

	utils.Handle(ctx, func() gin.H {
		return gin.H{"data": "Item deleted successfully"}
	}, err, http.StatusOK)
}
