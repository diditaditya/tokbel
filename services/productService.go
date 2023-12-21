package services

import (
	"errors"
	"tokbel/entity"
	"tokbel/services/data"
)

type ProductService struct {
	productRepo  data.ProductRepoInterface
	categoryRepo data.CategoryRepoInterface
}

func NewProductService(productRepo data.ProductRepoInterface, categoryRepo data.CategoryRepoInterface) *ProductService {
	return &ProductService{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

func (service *ProductService) Create(product *entity.Product) (*entity.Product, error) {
	_, err := service.categoryRepo.FindById(product.CategoryId)
	if err != nil {
		return nil, errors.New("category not found")
	}
	return service.productRepo.Create(product)
}

func (service *ProductService) FindAll() []entity.Product {
	return service.productRepo.FindAll()
}

func (service *ProductService) FindById(id int) (*entity.Product, error) {
	return service.productRepo.FindById(id)
}

func (service *ProductService) Update(product *entity.Product) (*entity.Product, error) {
	_, err := service.categoryRepo.FindById(product.CategoryId)
	if err != nil {
		return nil, errors.New("category not found")
	}
	return service.productRepo.Update(product)
}

func (service *ProductService) Delete(id int) error {
	return service.productRepo.Delete(id)
}
