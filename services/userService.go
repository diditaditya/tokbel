package services

import (
	"errors"
	"tokbel/auth"
	"tokbel/entity"
	"tokbel/services/data"
)

type UserService struct {
	userRepo data.UserRepoInterface
}

func NewUserService(userRepo data.UserRepoInterface) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (service *UserService) Register(user *entity.User) (*entity.User, error) {
	hashed, err := auth.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashed
	created, err := service.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (service *UserService) Login(email string, password string) (string, error) {
	found, err := service.userRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if !auth.CheckPassword(password, found.Password) {
		return "", errors.New("invalid password")
	}

	token, err := auth.CreateJWT(found.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *UserService) FindById(id int) (*entity.User, error) {
	found, err := service.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	return found, nil
}

func (service *UserService) FindByEmail(email string) (*entity.User, error) {
	found, err := service.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return found, nil
}

func (service *UserService) TopUp(id int, added int) (int, error) {
	found, err := service.userRepo.FindById(id)
	if err != nil {
		return 0, err
	}

	found.Balance = found.Balance + added
	updated, err := service.userRepo.Update(found)
	if err != nil {
		return 0, err
	}

	return updated.Balance, nil
}
