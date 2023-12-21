package repository

import (
	"errors"
	"log"
	"tokbel/auth"
	"tokbel/config"
	"tokbel/repository/category"
	"tokbel/repository/models"
	"tokbel/repository/product"
	"tokbel/repository/transaction"
	"tokbel/repository/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db          *gorm.DB
	User        *user.UserRepo
	Category    *category.CategoryRepo
	Product     *product.ProductRepo
	Transaction *transaction.TransactionHistoryRepo
}

func New(conf *config.DBConfig) *Repository {
	db := connectToDB(conf)
	repo := Repository{
		db:          db,
		User:        user.New(db),
		Category:    category.New(db),
		Product:     product.New(db),
		Transaction: transaction.New(db),
	}

	repo.prepareDB()

	return &repo
}

func connectToDB(conf *config.DBConfig) *gorm.DB {
	dsn := config.GetDBConfig().GetDBUrl()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database")
	}
	return db
}

func (repo *Repository) prepareDB() {
	repo.db.AutoMigrate(&models.User{}, &models.Category{},
		&models.Product{}, &models.TransactionHistory{})

	repo.seedAdmin()
}

func (repo *Repository) seedAdmin() {
	var admin models.User
	result := repo.db.Where("role = ?", "admin").First(&admin)

	if result.Error == nil {
		return
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Println("admin not found, creating one..")

		admin = models.User{
			FullName: "admin",
			Email:    "admin@email.com",
			Role:     "admin",
			Balance:  0,
		}

		defaultPassword := "password"
		hashed, err := auth.HashPassword(defaultPassword)
		if err != nil {
			log.Println("error hashing password for admin seed")
			return
		}

		admin.Password = hashed
		result := repo.db.Save(&admin)
		if result.Error != nil {
			log.Println("error seeding admin")
			return
		}

		log.Println("done seeding admin")
	}

	log.Println("something went wrong when seeding admin")
}
