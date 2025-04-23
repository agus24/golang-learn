package services

import (
	"golang_gin/app/ginapp_2/model"
	"golang_gin/app/repositories"
)

type CategoryService struct {
	Repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo}
}

func (self *CategoryService) GetAllCategories(search string, page int64, perPage int64) ([]model.Categories, error) {
	return self.Repo.GetAllCategories(search, page, perPage)
}

func (self *CategoryService) CreateCategory(name string) (*model.Categories, error) {
	result, err := self.Repo.CreateCategory(name)

	if err != nil {
		return nil, err
	}

	categoryId, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return self.Repo.GetCategoryById(categoryId)
}

func (self *CategoryService) GetCategory(id int64) (*model.Categories, error) {
	return self.Repo.GetCategoryById(id)
}

func (self *CategoryService) UpdateCategory(id int64, name string) (*model.Categories, error) {
	category, err := self.Repo.GetCategoryById(id)

	if err != nil {
		return nil, err
	}

	_, err = self.Repo.UpdateCategory(id, name)

	if err != nil {
		return nil, err
	}

	category, _ = self.Repo.GetCategoryById(id)

	return category, nil
}

func (self *CategoryService) DeleteCategory(category *model.Categories) error {
	return self.Repo.DeleteCategory(category.ID)
}
