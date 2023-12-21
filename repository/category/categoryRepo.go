package category

import (
	"errors"
	"log"
	"tokbel/entity"
	"tokbel/repository/models"

	"gorm.io/gorm"
)

type CategoryRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{db: db}
}

func convertModelToEntity(source *models.Category) entity.Category {
	target := entity.Category{}
	target.Id = int(source.ID)
	target.Type = source.Type
	target.SoldProductAmount = source.SoldProductAmount
	target.CreatedAt = source.CreatedAt
	target.UpdatedAt = source.UpdatedAt

	products := []entity.Product{}
	for _, prodModel := range source.Products {
		prodEntity := entity.Product{}
		prodEntity.Id = int(prodModel.ID)
		prodEntity.Title = prodModel.Title
		prodEntity.Price = prodModel.Price
		prodEntity.Stock = prodModel.Stock
		prodEntity.CategoryId = prodModel.CategoryID
		products = append(products, prodEntity)
	}
	target.Products = products

	return target
}

func convertEntityToModel(source *entity.Category) models.Category {
	target := models.Category{}
	target.ID = uint(source.Id)
	target.Type = source.Type
	target.SoldProductAmount = source.SoldProductAmount
	target.CreatedAt = source.CreatedAt
	target.UpdatedAt = source.UpdatedAt

	return target
}

func (repo *CategoryRepo) FindAll() []entity.Category {
	var found []models.Category
	result := repo.db.Preload("Products").Find(&found)
	if result.Error != nil {
		log.Println("error finding all categories")
		return []entity.Category{}
	}

	data := []entity.Category{}
	for _, raw := range found {
		data = append(data, convertModelToEntity(&raw))
	}
	return data
}

func (repo *CategoryRepo) FindById(id int) (*entity.Category, error) {
	var found models.Category
	result := repo.db.First(&found, id)
	if result.Error == nil {
		data := convertModelToEntity(&found)
		return &data, nil
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("data not found")
	}
	return nil, errors.New("error finding data")
}

func (repo *CategoryRepo) Create(category *entity.Category) (*entity.Category, error) {
	if category == nil {
		return nil, errors.New("data to save is required")
	}

	raw := convertEntityToModel(category)
	result := repo.db.Create(&raw)
	if result.Error != nil {
		return nil, result.Error
	}

	data := convertModelToEntity(&raw)
	return &data, nil
}

func (repo *CategoryRepo) Update(category *entity.Category) (*entity.Category, error) {
	var found models.Category
	result := repo.db.First(&found, category.Id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("data not found")
		}
		return nil, errors.New("something went wrong")
	}

	model := convertEntityToModel(category)
	result = repo.db.Save(&model)
	if result.Error != nil {
		return nil, errors.New("something went wrong")
	}

	data := convertModelToEntity(&model)
	return &data, nil
}

func (repo *CategoryRepo) Delete(id int) error {
	var found models.Category
	result := repo.db.First(&found, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("data not found")
		}
		return errors.New("something went wrong")
	}

	result = repo.db.Delete(&found)
	if result.Error != nil {
		return errors.New("something went wrong")
	}
	return nil
}
