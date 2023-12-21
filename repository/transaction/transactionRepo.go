package transaction

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

type TransactionHistoryRepo struct {
	db     *gorm.DB
	locker *utils.Locker
}

func New(db *gorm.DB, locker *utils.Locker) *TransactionHistoryRepo {
	return &TransactionHistoryRepo{db: db, locker: locker}
}

func convertModelToEntity(source *models.TransactionHistory) entity.TransactionHistory {
	target := entity.TransactionHistory{}
	target.Id = int(source.ID)
	target.UserId = source.UserID
	target.ProductId = source.ProductID
	target.Quantity = source.Quantity
	target.TotalPrice = source.TotalPrice
	target.Product = entity.Product{
		Id:         int(source.Product.ID),
		Title:      source.Product.Title,
		Price:      source.Product.Price,
		Stock:      source.Product.Stock,
		CategoryId: source.Product.CategoryID,
	}
	target.User = entity.User{
		Id:       int(source.User.ID),
		FullName: source.User.FullName,
		Email:    source.User.Email,
		Role:     source.User.Role,
		Balance:  source.User.Balance,
	}

	return target
}

func convertEntityToModel(source *entity.TransactionHistory) models.TransactionHistory {
	target := models.TransactionHistory{}
	target.ID = uint(source.Id)
	target.UserID = source.UserId
	target.ProductID = source.ProductId
	target.Quantity = source.Quantity
	target.TotalPrice = source.TotalPrice

	return target
}

func (repo *TransactionHistoryRepo) FindAll() []entity.TransactionHistory {
	var found []models.TransactionHistory
	result := repo.db.Preload("User").Preload("Product").Find(&found)
	if result.Error != nil {
		log.Println("error finding all categories")
		return []entity.TransactionHistory{}
	}

	data := []entity.TransactionHistory{}
	for _, raw := range found {
		data = append(data, convertModelToEntity(&raw))
	}
	return data
}

func (repo *TransactionHistoryRepo) FindByUserId(userId int) []entity.TransactionHistory {
	var found []models.TransactionHistory
	result := repo.db.Where("user_id = ?", userId).Preload("Product").Find(&found)
	if result.Error != nil {
		log.Println("error finding all categories")
		return []entity.TransactionHistory{}
	}

	data := []entity.TransactionHistory{}
	for _, raw := range found {
		data = append(data, convertModelToEntity(&raw))
	}
	return data
}

func (repo *TransactionHistoryRepo) Create(trx *entity.TransactionHistory, lock *data.Lock) (*entity.TransactionHistory, error) {
	if trx == nil {
		return nil, errors.New("data to save is required")
	}

	raw := convertEntityToModel(trx)

	var result *gorm.DB
	tx := repo.db
	if lock == nil {
		result = tx.Create(&raw)
	} else {
		tx = repo.locker.GetLock(lock)
		result = tx.Clauses(clause.Locking{Strength: "UPDATE"}).Create(&raw)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	var created models.TransactionHistory
	result = tx.Preload("Product").Preload("User").First(&created, raw.ID)
	if result.Error != nil {
		return nil, result.Error
	}

	data := convertModelToEntity(&created)
	return &data, nil
}
