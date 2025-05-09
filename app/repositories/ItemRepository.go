package repositories

import (
	"database/sql"
	"errors"

	"golang_gin/app/databases/model"
	. "golang_gin/app/databases/table"
	"golang_gin/app/utils"

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
		return nil, errors.New("Item not found.")
	}

	return &results[0], err
}

func (self *ItemRepository) getDefaultQuery() SelectStatement {
	return SELECT(
		Items.AllColumns,
		SubCategories.AllColumns,
		Categories.AllColumns,
	).FROM(Items.
		INNER_JOIN(SubCategories, Items.SubCategoryID.EQ(SubCategories.ID)).
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

func (self *ItemRepository) Create(name string, price int, subCategoryID int64) (*Item, error) {
	utils.StartTransaction(self.db)
	stmt := Items.INSERT(Items.Name, Items.Price, Items.SubCategoryID).
		VALUES(name, price, subCategoryID)

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

func (self *ItemRepository) Update(id int64, name string, price int, subCategoryID int64) (*Item, error) {
	utils.StartTransaction(self.db)
	stmt := Items.UPDATE(Items.Name, Items.Price, Items.SubCategoryID).
		SET(name, price, subCategoryID).
		WHERE(Items.ID.EQ(Int64(id)))

	_, err := stmt.Exec(self.db)

	if err != nil {
		return nil, err
	}

	return self.GetById(id)
}

func (self *ItemRepository) GetItemByIds(ids []int64) ([]Item, error) {
	exprs := make([]Expression, len(ids))
	for i, id := range ids {
		exprs[i] = Int64(id)
	}

	stmt := SELECT(Items.AllColumns).FROM(Items).WHERE(Items.ID.IN(exprs...))

	return self.getMultiple(stmt)
}

func (self *ItemRepository) Delete(id int64) error {
	utils.StartTransaction(self.db)
	stmt := Items.DELETE().WHERE(Items.ID.EQ(Int64(id)))

	_, err := stmt.Exec(self.db)

	return err
}
