package data

import "tokbel/entity"

type UserRepoInterface interface {
	FindAll() []entity.User
	FindById(id int) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Create(user *entity.User) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
}
