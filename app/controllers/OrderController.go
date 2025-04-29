package controllers

import (
	"database/sql"
	"golang_gin/app/repositories"
	"golang_gin/app/serializers"
	"golang_gin/app/services"
	"golang_gin/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	service services.OrderService
}

func NewOrderController(db *sql.DB) OrderController {
	repo := repositories.NewOrderRepository(db)
	return OrderController{services.NewOrderService(repo)}
}

func (self OrderController) GetAll(ctx *gin.Context) {
	search := ctx.DefaultQuery("search", "")
	pageQuery := ctx.DefaultQuery("page", "1")
	perPageQuery := ctx.DefaultQuery("per_page", "")

	page, perPage, err := utils.ParsePageAndPerPage(pageQuery, perPageQuery)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid page or per_page value", "error_code": "invalid_page_or_per_page"})
		return
	}

	orders, err := self.service.GetAll(search, *page, *perPage)

	utils.Handle(ctx, func() gin.H {
		return gin.H{"data": serializers.Orders(orders), "meta": serializers.Pagination(page, perPage)}
	}, err, http.StatusOK)
}
