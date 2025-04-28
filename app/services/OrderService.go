package services

import (
	"golang_gin/app/databases/model"
	"golang_gin/app/repositories"
	"time"
)

type OrderService struct {
	Repo repositories.OrderRepository
}

func NewOrderService(repository repositories.OrderRepository) OrderService {
	return OrderService{Repo: repository}
}

func (self *OrderService) GetAll(search string, page int64, perPage int64) ([]model.Orders, error) {
	return self.Repo.GetAll(search, page, perPage)
}

func (self *OrderService) GetOrderById(id int64) (*model.Orders, error) {
	return self.Repo.GetOrderById(id)
}

func (self *OrderService) CreateOrder(date time.Time, orderNumber string, grandTotal int, customerName string) (*model.Orders, error) {
	return self.Repo.CreateOrder(date, orderNumber, grandTotal, customerName)
}
