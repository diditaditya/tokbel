package product

import (
	"errors"
	"log"
	"tokbel/entity"
	"tokbel/repository/models"
	"tokbel/repository/utils"
	"tokbel/services/data"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepo struct {
	db     *gorm.DB
	locker *utils.Locker
}

func New(db *gorm.DB, locker *utils.Locker) *ProductRepo {
	return &ProductRepo{db: db, locker: locker}
}

func convertModelToEntity(source *models.Product) entity.Product {
	target := entity.Product{}
	target.Id = int(source.ID)
	target.Title = source.Title
	target.Price = source.Price
	target.Stock = source.Stock
	target.CreatedAt = source.CreatedAt
	target.CategoryId = source.CategoryID
	target.Category = entity.Category{
		Id:                int(source.Category.ID),
		Type:              source.Category.Type,
		SoldProductAmount: source.Category.SoldProductAmount,
	}

	return target
}

func convertEntityToModel(source *entity.Product) models.Product {
	target := models.Product{}
	target.ID = uint(source.Id)
	target.Title = source.Title
	target.Price = source.Price
	target.Stock = source.Stock
	target.CategoryID = source.CategoryId
	target.CreatedAt = source.CreatedAt

	return target
}

func (repo *ProductRepo) FindAll() []entity.Product {
	var found []models.Product
	result := repo.db.Find(&found)
	if result.Error != nil {
		log.Println("error finding all categories")
		return []entity.Product{}
	}

	data := []entity.Product{}
	for _, raw := range found {
		data = append(data, convertModelToEntity(&raw))
	}
	return data
}

func (repo *ProductRepo) FindById(id int, lock *data.Lock) (*entity.Product, error) {
	var found models.Product
	var result *gorm.DB
	tx := repo.db
	if lock == nil {
		result = tx.First(&found, id)
	} else {
		tx = repo.locker.GetLock(lock)
		result = tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&found, id)
	}
	if result.Error == nil {
		data := convertModelToEntity(&found)
		return &data, nil
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("data not found")
	}
	return nil, errors.New("error finding data")
}

func (repo *ProductRepo) Create(product *entity.Product) (*entity.Product, error) {
	if product == nil {
		return nil, errors.New("data to create is required")
	}

	raw := convertEntityToModel(product)
	result := repo.db.Create(&raw)
	if result.Error != nil {
		return nil, result.Error
	}

	data := convertModelToEntity(&raw)
	return &data, nil
}

func (repo *ProductRepo) Update(product *entity.Product, lock *data.Lock) (*entity.Product, error) {
	var found models.Product
	var result *gorm.DB
	tx := repo.db
	if lock == nil {
		result = tx.First(&found, product.Id)
	} else {
		tx = repo.locker.GetLock(lock)
		result = tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&found, product.Id)
	}
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("data not found")
		}
		return nil, errors.New("something went wrong")
	}

	model := convertEntityToModel(product)
	result = tx.Save(&model)
	if result.Error != nil {
		return nil, errors.New("something went wrong")
	}

	data := convertModelToEntity(&model)
	return &data, nil
}

func (repo *ProductRepo) Delete(id int) error {
	var found models.Product
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
