package repositories

import (
	"database/sql"
	"errors"

	"golang_gin/app/ginapp_2/model"
	. "golang_gin/app/ginapp_2/table"

	. "github.com/go-jet/jet/v2/mysql"
)

type ItemRepository struct {
	db *sql.DB
}

type Item struct {
	model.Items

	SubCategory *SubCategory
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{db}
}

func (self *ItemRepository) getMultiple(stmt SelectStatement) ([]Item, error) {
	var results []Item

	err := stmt.Query(self.db, &results)

	return results, err
}

func (self *ItemRepository) getSingle(stmt SelectStatement) (*Item, error) {
	results, err := self.getMultiple(stmt)

	if len(results) == 0 {
		return nil, errors.New("Sub Category not found.")
	}

	return &results[0], err
}

func (self *ItemRepository) getDefaultQuery() SelectStatement {
	return SELECT(
		Items.AllColumns,
		SubCategories.AllColumns,
		Categories.AllColumns,
	).FROM(Items.
		INNER_JOIN(SubCategories, Items.ID.EQ(SubCategories.ID)).
		INNER_JOIN(Categories, SubCategories.CategoryID.EQ(Categories.ID)),
	)
}

func (self *ItemRepository) GetAll(search string, page int64, perPage int64) ([]Item, error) {
	stmt := self.getDefaultQuery()

	if search != "" {
		stmt = stmt.WHERE(
			Items.Name.LIKE(String("%" + search + "%")).
				OR(SubCategories.Name.LIKE(String("%" + search + "%"))),
		)
	}

	if page > 0 && perPage > 0 {
		stmt = stmt.LIMIT(perPage).OFFSET((page - 1) * perPage)
	}

	return self.getMultiple(stmt)
}

func (self *ItemRepository) GetById(id int64) (*Item, error) {
	stmt := self.getDefaultQuery().
		WHERE(Items.ID.EQ(Int64(id)))

	return self.getSingle(stmt)
}

func (self *ItemRepository) Create(name string, price int, subCategoryId int64) (*Item, error) {
	stmt := Items.INSERT(Items.Name, Items.Price, Items.SubCategoryID).VALUES(name, price, subCategoryId)

	result, err := stmt.Exec(self.db)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return self.GetById(id)
}

func (self *ItemRepository) Update(id int64, name string, price int, subCategoryId int64) (*Item, error) {
	stmt := Items.UPDATE(Items.Name, Items.Price, Items.SubCategoryID).SET(name, price, subCategoryId).WHERE(Items.ID.EQ(Int64(id)))

	_, err := stmt.Exec(self.db)

	if err != nil {
		return nil, err
	}

	return self.GetById(id)
}

func (self *ItemRepository) Delete(id int64) error {
	stmt := Items.DELETE().WHERE(Items.ID.EQ(Int64(id)))

	_, err := stmt.Exec(self.db)

	return err
}
