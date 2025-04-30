package dtos

import "time"

type OrderDTO struct {
	CustomerName string
	Date         time.Time
	OrderNumber  string
	GrandTotal   int
	Details      []OrderDetailDTO
}

type OrderDetailDTO struct {
	ItemID   int64
	Quantity int
	Price    int32
}
