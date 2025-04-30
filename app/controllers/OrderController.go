package controllers

import (
	"database/sql"
	"errors"
	"golang_gin/app/dtos"
	"golang_gin/app/repositories"
	"golang_gin/app/requests"
	"golang_gin/app/serializers"
	"golang_gin/app/services"
	"golang_gin/app/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	service        services.OrderService
	itemRepository repositories.ItemRepository
}

func NewOrderController(db *sql.DB) OrderController {
	repo := repositories.NewOrderRepository(db)
	return OrderController{
		service:        services.NewOrderService(repo),
		itemRepository: *repositories.NewItemRepository(db),
	}
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

func (self OrderController) Create(ctx *gin.Context) {
	raw, _ := ctx.Get("validated")
	input := raw.(requests.CreateOrderRequest)

	date := *self.validateDate(ctx, input)
	formattedItem, err := self.validateItem(ctx, input)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Items not found", "error_code": "items_not_found"})
		return
	}

	details := make([]dtos.OrderDetailDTO, len(input.Details))

	for key, detail := range input.Details {
		details[key] = dtos.OrderDetailDTO{
			ItemID:   detail.ItemID,
			Quantity: detail.Quantity,
		}
	}

	order, err := self.service.CreateOrder(dtos.OrderDTO{
		CustomerName: input.CustomerName,
		Date:         date,
		Details:      details,
	}, formattedItem)

	utils.Handle(ctx, func() gin.H {
		return gin.H{"data": serializers.Order(order)}
	}, err, http.StatusOK)
}

func (self OrderController) Show(ctx *gin.Context) {
	id := utils.ParseToInt(ctx, ctx.Param("id"))

	order, err := self.service.GetById(*id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": serializers.Order(order)})
}

func (self OrderController) validateDate(ctx *gin.Context, input requests.CreateOrderRequest) *time.Time {
	date, err := utils.ParseDate(input.Date)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid date format", "error_code": "invalid_date_format"})
		return nil
	}

	return &date
}

func (self OrderController) validateItem(ctx *gin.Context, input requests.CreateOrderRequest) (map[int64]repositories.Item, error) {
	itemIds := make([]int64, len(input.Details))
	formattedItem := make(map[int64]repositories.Item)

	for key, value := range input.Details {
		itemIds[key] = value.ItemID
	}

	items, err := self.itemRepository.GetItemByIds(itemIds)

	if err != nil {
		return nil, err
	}

	if len(itemIds) != len(items) {
		return nil, errors.New("Item Not Found")
	}

	for _, item := range items {
		formattedItem[item.ID] = item
	}

	return formattedItem, nil
}
