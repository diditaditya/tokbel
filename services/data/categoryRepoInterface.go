package data

import "tokbel/entity"

type CategoryRepoInterface interface {
	FindAll() []entity.Category
	FindById(id int, lock *Lock) (*entity.Category, error)
	Create(category *entity.Category) (*entity.Category, error)
	Update(category *entity.Category, lock *Lock) (*entity.Category, error)
	Delete(categoryId int) error
}
