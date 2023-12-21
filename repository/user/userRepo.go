package user

import (
	"errors"
	"log"
	"tokbel/entity"
	"tokbel/repository/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func convertModelToEntity(source *models.User) entity.User {
	target := entity.User{}
	target.Id = int(source.ID)
	target.FullName = source.FullName
	target.Email = source.Email
	target.Password = source.Password
	target.Role = source.Role
	target.Balance = source.Balance
	target.CreatedAt = source.CreatedAt
	target.UpdatedAt = source.UpdatedAt

	return target
}

func convertEntityToModel(source *entity.User) models.User {
	target := models.User{}
	target.ID = uint(source.Id)
	target.FullName = source.FullName
	target.Email = source.Email
	target.Password = source.Password
	target.Role = source.Role
	target.Balance = source.Balance
	target.CreatedAt = source.CreatedAt
	target.UpdatedAt = source.UpdatedAt

	return target
}

func (repo *UserRepo) FindAll() []entity.User {
	var found []models.User
	result := repo.db.Find(&found)
	if result.Error != nil {
		log.Println("error finding all users")
		return []entity.User{}
	}

	data := []entity.User{}
	for _, raw := range found {
		data = append(data, convertModelToEntity(&raw))
	}
	return data
}

func (repo *UserRepo) FindById(id int) (*entity.User, error) {
	var found models.User
	result := repo.db.First(&found, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("data not found")
		}
		return nil, errors.New("something went wrong")
	}

	data := convertModelToEntity(&found)
	return &data, nil
}

func (repo *UserRepo) FindByEmail(email string) (*entity.User, error) {
	var found models.User
	result := repo.db.Where("email = ?", email).First(&found)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("data not found")
		}
		return nil, errors.New("something went wrong")
	}

	data := convertModelToEntity(&found)
	return &data, nil
}

func (repo *UserRepo) Create(user *entity.User) (*entity.User, error) {
	model := convertEntityToModel(user)
	result := repo.db.Save(&model)
	if result.Error != nil {
		return nil, result.Error
	}

	data := convertModelToEntity(&model)
	return &data, nil
}

func (repo *UserRepo) Update(user *entity.User) (*entity.User, error) {
	var found models.User
	result := repo.db.First(&found, user.Id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("data not found")
		}
		return nil, errors.New("something went wrong")
	}

	model := convertEntityToModel(user)
	result = repo.db.Save(&model)
	if result.Error != nil {
		return nil, errors.New("something went wrong")
	}

	data := convertModelToEntity(&model)
	return &data, nil
}
