package repositories

import (
	"database/sql"
	"errors"

	. "github.com/go-jet/jet/v2/mysql"
	"golang_gin/app/ginapp_2/model"
	. "golang_gin/app/ginapp_2/table"
)

type SubCategory struct {
	model.SubCategories

	Category *model.Categories
}

type SubCategoryRepository struct {
	db *sql.DB
}

func NewSubCategoryRepository(db *sql.DB) *SubCategoryRepository {
	return &SubCategoryRepository{db}
}

func (self *SubCategoryRepository) getDefaultSelect() SelectStatement {
	return SELECT(
		SubCategories.AllColumns,
		Categories.AllColumns,
	).FROM(SubCategories.
		INNER_JOIN(Categories, SubCategories.CategoryID.EQ(Categories.ID)),
	)
}

func (self *SubCategoryRepository) getMultiple(stmt SelectStatement) ([]SubCategory, error) {
	var results []SubCategory

	err := stmt.Query(self.db, &results)

	return results, err
}

func (self *SubCategoryRepository) getSingle(stmt SelectStatement) (*SubCategory, error) {
	results, err := self.getMultiple(stmt)

	if len(results) == 0 {
		return nil, errors.New("Sub Category not found.")
	}

	return &results[0], err
}

func (self *SubCategoryRepository) GetAllSubCategories(search string, page int64, perPage int64) ([]SubCategory, error) {
	stmt := self.getDefaultSelect()

	if search != "" {
		searchValue := String("%" + search + "%")
		stmt = stmt.WHERE(
			SubCategories.Name.LIKE(searchValue).
				OR(Categories.Name.LIKE(searchValue)),
		)
	}

	if page > 0 && perPage > 0 {
		stmt = stmt.LIMIT(perPage).OFFSET((page - 1) * perPage)
	}

	return self.getMultiple(stmt)
}

func (self *SubCategoryRepository) GetSubCategoryById(id int64) (*SubCategory, error) {
	stmt := self.getDefaultSelect()
	stmt = stmt.WHERE(SubCategories.ID.EQ(Int64(id)))

	return self.getSingle(stmt)
}

func (self *SubCategoryRepository) GetSubCategoryByName(name string) (*SubCategory, error) {
	stmt := self.getDefaultSelect()
	stmt = stmt.WHERE(SubCategories.Name.EQ(String(name)))

	return self.getSingle(stmt)
}

func (self *SubCategoryRepository) CreateSubCategory(name string, categoryID int64) (*SubCategory, error) {
	stmt := SubCategories.INSERT(SubCategories.Name, SubCategories.CategoryID).
		VALUES(name, categoryID)

	result, err := stmt.Exec(self.db)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return self.GetSubCategoryById(id)
}

func (self *SubCategoryRepository) UpdateSubCategory(id int64, name string, categoryID int64) (sql.Result, error) {
	stmt := SubCategories.UPDATE(
		SubCategories.Name,
		SubCategories.CategoryID,
	).SET(name, categoryID).WHERE(SubCategories.ID.EQ(Int64(id)))

	return stmt.Exec(self.db)
}

func (self *SubCategoryRepository) DeleteSubCategory(id int64) error {
	stmt := SubCategories.DELETE().
		WHERE(SubCategories.ID.EQ(Int64(id)))

	_, err := stmt.Exec(self.db)

	return err
}
