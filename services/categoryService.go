package services

import (
	"tokbel/entity"
	"tokbel/services/data"
)

type CategoryService struct {
	categoryRepo data.CategoryRepoInterface
}

func NewCategoryService(repo data.CategoryRepoInterface) *CategoryService {
	return &CategoryService{
		categoryRepo: repo,
	}
}

func (service *CategoryService) Create(category *entity.Category) (*entity.Category, error) {
	return service.categoryRepo.Create(category)
}

func (service *CategoryService) FindAll() []entity.Category {
	return service.categoryRepo.FindAll()
}

func (service *CategoryService) UpdateType(id int, categoryType string) (*entity.Category, error) {
	found, err := service.categoryRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	found.Type = categoryType
	return service.categoryRepo.Update(found)
}

func (service *CategoryService) Delete(id int) error {
	return service.categoryRepo.Delete(id)
}
