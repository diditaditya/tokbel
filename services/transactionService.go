package services

import (
	"errors"
	"fmt"
	"tokbel/entity"
	"tokbel/services/data"
)

type TransactionService struct {
	trxRepo      data.TransactionRepoInterface
	productRepo  data.ProductRepoInterface
	userRepo     data.UserRepoInterface
	categoryRepo data.CategoryRepoInterface
}

func NewTransactionService(
	trxRepo data.TransactionRepoInterface,
	productRepo data.ProductRepoInterface,
	userRepo data.UserRepoInterface,
	categoryRepo data.CategoryRepoInterface) *TransactionService {

	return &TransactionService{
		trxRepo:      trxRepo,
		productRepo:  productRepo,
		userRepo:     userRepo,
		categoryRepo: categoryRepo,
	}
}

func (service *TransactionService) FindAll() []entity.TransactionHistory {
	return service.trxRepo.FindAll()
}

func (service *TransactionService) FindByUserId(userId int) []entity.TransactionHistory {
	return service.trxRepo.FindByUserId(userId)
}

func (service *TransactionService) Create(trx *entity.TransactionHistory) (*entity.TransactionHistory, error) {
	// NOTE: this process is prone to race condition!
	// NOTE: use locking mechanism and db transaction!
	// NOTE: the following processes can be run simultaneously

	// check if product exists
	product, err := service.productRepo.FindById(trx.ProductId)
	if err != nil {
		fmt.Println("-->>> product not found")
		return nil, err
	}
	// check if quantity less than stock
	if product.Stock < trx.Quantity {
		fmt.Println("-->>> stock is less than quantity", product.Stock, " vs ", trx.Quantity)
		return nil, errors.New("not enough stock")
	}

	// calculate total price
	trx.TotalPrice = product.Price * trx.Quantity

	// check if user balance is greater than the total price
	user, err := service.userRepo.FindById(trx.UserId)
	if err != nil {
		fmt.Println("-->>> user not found")
		return nil, err
	}
	if user.Balance < trx.TotalPrice {
		fmt.Println("-->>> user do not have enough balance ", user.Balance, " vs total price: ", trx.TotalPrice)
		return nil, errors.New("not enough balance")
	}

	// update stock of product
	product.Stock -= trx.Quantity
	_, err = service.productRepo.Update(product)
	if err != nil {
		fmt.Println("-->>> error updating product")
		return nil, err
	}
	// update balance of user
	user.Balance -= trx.TotalPrice
	_, err = service.userRepo.Update(user)
	if err != nil {
		fmt.Println("-->>> error updating user")
		return nil, err
	}
	// update sold amount of category
	category, err := service.categoryRepo.FindById(product.CategoryId)
	if err != nil {
		fmt.Println("-->>> category not found")
		return nil, err
	}
	category.SoldProductAmount += trx.Quantity
	_, err = service.categoryRepo.Update(category)
	if err != nil {
		fmt.Println("-->>> error updating category sold product amount")
		return nil, err
	}

	created, err := service.trxRepo.Create(trx)
	if err != nil {
		fmt.Println("-->>> error creating transaction history")
		return nil, err
	}

	return created, nil
}
