package repositories

import (
	"database/sql"
	"errors"
	"golang_gin/app/databases/model"
	. "golang_gin/app/databases/table"
	"golang_gin/app/dtos"
	"golang_gin/app/utils"

	. "github.com/go-jet/jet/v2/mysql"
)

type OrderRepository struct {
	Db *sql.DB
}

type Order struct {
	model.Orders
	Details []OrderDetail
}

type OrderDetail struct {
	model.OrderDetails
	ItemName string
}

type OrderRow struct {
	model.Orders
	model.OrderDetails
	model.Items
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return OrderRepository{Db: db}
}

func (self *OrderRepository) getMultiple(stmt SelectStatement) ([]Order, error) {
	var results []OrderRow

	err := stmt.Query(self.Db, &results)

	return self.mapFromRow(results), err
}

func (self *OrderRepository) mapFromRow(rows []OrderRow) (orders []Order) {
	var orderMap = make(map[int64]*Order)

	for _, row := range rows {
		order, exists := orderMap[row.Orders.ID]
		if !exists {
			order = &Order{Orders: row.Orders}
			orderMap[row.Orders.ID] = order
		}

		orderDetail := OrderDetail{
			OrderDetails: row.OrderDetails,
			ItemName:     row.Items.Name,
		}

		order.Details = append(order.Details, orderDetail)
	}

	for _, order := range orderMap {
		orders = append(orders, *order)
	}

	return orders
}

func (self *OrderRepository) getSingle(stmt SelectStatement) (*Order, error) {
	results, err := self.getMultiple(stmt)

	if len(results) == 0 {
		return nil, errors.New("Order not found.")
	}

	return &results[0], err
}

func (self *OrderRepository) GetAll(search string, page int64, perPage int64) ([]Order, error) {
	stmt := SELECT(Orders.AllColumns).FROM(Orders)

	if search != "" {
		stmt = stmt.WHERE(
			Orders.CustomerName.LIKE(String("%" + search + "%")).
				OR(Orders.OrderNumber.LIKE(String("%" + search + "%"))),
		)
	}

	if page > 0 && perPage > 0 {
		stmt = stmt.LIMIT(perPage).OFFSET((page - 1) * perPage)
	}

	return self.getMultiple(stmt)
}

func (self *OrderRepository) GetOrderByIdWithDetail(id int64) (*Order, error) {
	stmt := SELECT(Orders.AllColumns, OrderDetails.AllColumns, Items.Name).FROM(
		Orders.INNER_JOIN(OrderDetails, OrderDetails.OrderID.EQ(Orders.ID)).
			INNER_JOIN(Items, Items.ID.EQ(OrderDetails.ItemID)),
	).WHERE(Orders.ID.EQ(Int(id)))

	return self.getSingle(stmt)
}

func (self *OrderRepository) GetOrderById(id int64) (*Order, error) {
	stmt := SELECT(Orders.AllColumns).FROM(Orders).WHERE(Orders.ID.EQ(Int(id)))

	return self.getSingle(stmt)
}

func (self *OrderRepository) CreateOrder(input dtos.OrderDTO) (*Order, error) {
	utils.StartTransaction(self.Db)

	stmt := Orders.INSERT(Orders.Date, Orders.OrderNumber, Orders.GrandTotal, Orders.CustomerName).
		VALUES(input.Date, input.OrderNumber, input.GrandTotal, input.CustomerName)

	result, err := stmt.Exec(self.Db)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	detailStmt := OrderDetails.INSERT(OrderDetails.OrderID, OrderDetails.ItemID, OrderDetails.Quantity, OrderDetails.Price)

	for _, detail := range input.Details {
		detailStmt = detailStmt.VALUES(id, detail.ItemID, detail.Quantity, detail.Price)
	}

	result, err = detailStmt.Exec(self.Db)

	if err != nil {
		return nil, err
	}

	return self.GetOrderById(id)
}
