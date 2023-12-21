package data

import "tokbel/entity"

type ProductRepoInterface interface {
	FindAll() []entity.Product
	FindById(id int) (*entity.Product, error)
	Create(product *entity.Product) (*entity.Product, error)
	Update(product *entity.Product) (*entity.Product, error)
	Delete(productId int) error
}
