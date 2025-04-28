package repositories

import (
	"database/sql"
	"errors"
	"golang_gin/app/databases/model"
	. "golang_gin/app/databases/table"
	"golang_gin/utils"
	"time"

	. "github.com/go-jet/jet/v2/mysql"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return OrderRepository{Db: db}
}

func (self *OrderRepository) getMultiple(stmt SelectStatement) ([]model.Orders, error) {
	var results []model.Orders

	err := stmt.Query(self.Db, &results)

	return results, err
}

func (self *OrderRepository) getSingle(stmt SelectStatement) (*model.Orders, error) {
	results, err := self.getMultiple(stmt)

	if len(results) == 0 {
		return nil, errors.New("Order not found.")
	}

	return &results[0], err
}

func (self *OrderRepository) GetAll(search string, page int64, perPage int64) ([]model.Orders, error) {
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

func (self *OrderRepository) GetOrderById(id int64) (*model.Orders, error) {
	stmt := SELECT(Orders.AllColumns).FROM(Orders).WHERE(Orders.ID.EQ(Int(id)))

	return self.getSingle(stmt)
}

func (self *OrderRepository) CreateOrder(date time.Time, orderNumber string, grandTotal int, customerName string) (*model.Orders, error) {
	utils.StartTransaction(self.Db)

	stmt := Orders.INSERT(Orders.Date, Orders.OrderNumber, Orders.GrandTotal, Orders.CustomerName).
		VALUES(date, orderNumber, grandTotal, customerName)

	result, err := stmt.Exec(self.Db)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return self.GetOrderById(id)
}

func (self *OrderRepository) UpdateGrandTotal(id int64, grandTotal int) error {
	stmt := Orders.UPDATE(Orders.GrandTotal).SET(grandTotal).WHERE(Orders.ID.EQ(Int(id)))

	_, err := stmt.Exec(self.Db)

	return err
}

func (self *OrderRepository) CreateDetail(orderId int64, itemId int64, quantity int, price int) error {
	utils.StartTransaction(self.Db)
	stmt := OrderDetails.INSERT(OrderDetails.OrderID, OrderDetails.ItemID, OrderDetails.Quantity, OrderDetails.Price).
		VALUES(orderId, itemId, quantity, price)

	_, err := stmt.Exec(self.Db)

	return err
}
