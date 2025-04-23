package repositories

import (
	"database/sql"
	"errors"
	. "github.com/go-jet/jet/v2/mysql"
	"golang_gin/app/ginapp_2/model"
	. "golang_gin/app/ginapp_2/table"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db}
}

func (self *CategoryRepository) getMultipleCategories(stmt SelectStatement) ([]model.Categories, error) {
	var results []model.Categories

	err := stmt.Query(self.db, &results)

	return results, err
}

func (self *CategoryRepository) getSingleCategory(stmt SelectStatement) (*model.Categories, error) {
	var results []model.Categories

	err := stmt.Query(self.db, &results)

	if len(results) == 0 {
		return nil, errors.New("Category not found.")
	}

	return &results[0], err
}

func (self *CategoryRepository) GetAllCategories(search string, page int64, perPage int64) ([]model.Categories, error) {
	stmt := SELECT(Categories.AllColumns).FROM(Categories)

	if search != "" {
		stmt = stmt.WHERE(Categories.Name.LIKE(String("%" + search + "%")))
	}

	if page > 0 && perPage > 0 {
		stmt = stmt.LIMIT(perPage).OFFSET((page - 1) * perPage)
	}

	return self.getMultipleCategories(stmt)
}

func (self *CategoryRepository) GetCategoryById(id int64) (*model.Categories, error) {
	stmt := SELECT(
		Categories.AllColumns,
	).FROM(Categories).
		WHERE(Categories.ID.EQ(Int64(id)))

	return self.getSingleCategory(stmt)
}

func (self *CategoryRepository) GetCategoryByName(name string, exclude int64) (*model.Categories, error) {
	stmt := SELECT(
		Categories.AllColumns,
	).FROM(Categories).
		WHERE(Categories.Name.EQ(String(name)))

	if exclude > 0 {
		stmt = stmt.WHERE(Categories.ID.NOT_EQ(Int64(exclude)))
	}

	return self.getSingleCategory(stmt)
}

func (self *CategoryRepository) CreateCategory(name string) (sql.Result, error) {
	stmt := Categories.INSERT(Categories.Name).VALUES(name)

	return stmt.Exec(self.db)
}

func (self *CategoryRepository) UpdateCategory(id int64, name string) (sql.Result, error) {
	stmt := Categories.UPDATE(Categories.Name).SET(name).
		WHERE(Categories.ID.EQ(Int64(id)))

	return stmt.Exec(self.db)
}

func (self *CategoryRepository) DeleteCategory(id int64) error {
	stmt := Categories.DELETE().
		WHERE(Categories.ID.EQ(Int64(id)))

	_, err := stmt.Exec(self.db)

	return err
}
