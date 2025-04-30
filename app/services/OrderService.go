package services

import (
	"fmt"
	"golang_gin/app/dtos"
	"golang_gin/app/repositories"
	"math/rand"
	"time"
)

type OrderService struct {
	Repo repositories.OrderRepository
}

func NewOrderService(repository repositories.OrderRepository) OrderService {
	return OrderService{Repo: repository}
}

func (self *OrderService) GetAll(search string, page int64, perPage int64) ([]repositories.Order, error) {
	return self.Repo.GetAll(search, page, perPage)
}

func (self *OrderService) GetOrderById(id int64) (*repositories.Order, error) {
	return self.Repo.GetOrderById(id)
}

func (self *OrderService) CreateOrder(input dtos.OrderDTO, formattedItem map[int64]repositories.Item) (*repositories.Order, error) {
	for key, detail := range input.Details {
		item := formattedItem[detail.ItemID]

		input.GrandTotal += int(item.Price) * detail.Quantity
		input.Details[key].Price = item.Price
	}

	input.OrderNumber = self.GenerateOrderNumber()

	return self.Repo.CreateOrder(input)
}

func (self *OrderService) GetById(id int64) (*repositories.Order, error) {
	return self.Repo.GetOrderByIdWithDetail(id)
}

func (self *OrderService) GenerateOrderNumber() string {
	now := time.Now()

	datePart := now.Format("20060102")

	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var randomPart string
	for i := 0; i < 5; i++ {
		randomPart += string(charset[rand.Intn(len(charset))])
	}

	return fmt.Sprintf("ORD-%s-%s", datePart, randomPart)
}
