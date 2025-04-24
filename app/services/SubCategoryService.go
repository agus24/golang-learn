package services

import (
	"golang_gin/app/repositories"
	"golang_gin/app/requests"
)

type SubCategoryService struct {
	Repo *repositories.SubCategoryRepository
}

func NewSubCategoryService(repo *repositories.SubCategoryRepository) *SubCategoryService {
	return &SubCategoryService{Repo: repo}
}

func (self *SubCategoryService) GetAllSubCategories(search string, page int64, perPage int64) ([]repositories.SubCategory, error) {
	return self.Repo.GetAllSubCategories(search, page, perPage)
}

func (self *SubCategoryService) GetSubCategoryById(id int64) (*repositories.SubCategory, error) {
	return self.Repo.GetSubCategoryById(id)
}

func (self *SubCategoryService) CreateSubCategory(body requests.SubCategoryCreateRequest) (*repositories.SubCategory, error) {
	return self.Repo.CreateSubCategory(body.Name, body.CategoryID)
}

func (self *SubCategoryService) UpdateSubCategory(id int64, body requests.SubCategoryUpdateRequest) (*repositories.SubCategory, error) {
	_, err := self.Repo.UpdateSubCategory(id, body.Name, body.CategoryID)

	if err != nil {
		return nil, err
	}

	return self.Repo.GetSubCategoryById(id)
}

func (self *SubCategoryService) DeleteSubCategory(id int64) error {
	err := self.Repo.DeleteSubCategory(id)

	return err
}
