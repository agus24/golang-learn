package services

import (
	"golang_gin/app/repositories"
	"golang_gin/app/requests"
)

type ItemService struct {
	ItemRepository *repositories.ItemRepository
}

func NewItemService(itemRepository *repositories.ItemRepository) *ItemService {
	return &ItemService{
		ItemRepository: itemRepository,
	}
}

func (self *ItemService) GetAll(search string, page int64, perPage int64) ([]repositories.Item, error) {
	return self.ItemRepository.GetAll(search, page, perPage)
}

func (self *ItemService) GetById(id int64) (*repositories.Item, error) {
	return self.ItemRepository.GetById(id)
}

func (self *ItemService) Create(input requests.ItemCreateRequest) (*repositories.Item, error) {
	return self.ItemRepository.Create(input.Name, input.Price, input.SubCategoryID)
}

func (self *ItemService) Update(id int64, input requests.ItemUpdateRequest) (*repositories.Item, error) {
	return self.ItemRepository.Update(id, input.Name, input.Price, input.SubCategoryID)
}

func (self *ItemService) Delete(id int64) error {
	return self.ItemRepository.Delete(id)
}
