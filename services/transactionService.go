package services

import (
	"errors"
	"tokbel/entity"
	"tokbel/services/data"
)

type TransactionService struct {
	locker       data.LockerInterface
	trxRepo      data.TransactionRepoInterface
	productRepo  data.ProductRepoInterface
	userRepo     data.UserRepoInterface
	categoryRepo data.CategoryRepoInterface
}

func NewTransactionService(
	locker data.LockerInterface,
	trxRepo data.TransactionRepoInterface,
	productRepo data.ProductRepoInterface,
	userRepo data.UserRepoInterface,
	categoryRepo data.CategoryRepoInterface) *TransactionService {

	return &TransactionService{
		locker:       locker,
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

	lock := service.locker.StartLock()

	// check if product exists
	product, err := service.productRepo.FindById(trx.ProductId, lock)
	if err != nil {
		service.locker.Abort(lock)
		return nil, err
	}
	// check if quantity less than stock
	if product.Stock < trx.Quantity {
		service.locker.Abort(lock)
		return nil, errors.New("not enough stock")
	}

	// calculate total price
	trx.TotalPrice = product.Price * trx.Quantity

	// check if user balance is greater than the total price
	user, err := service.userRepo.FindById(trx.UserId, lock)
	if err != nil {
		service.locker.Abort(lock)
		return nil, err
	}
	if user.Balance < trx.TotalPrice {
		service.locker.Abort(lock)
		return nil, errors.New("not enough balance")
	}

	// update stock of product
	product.Stock -= trx.Quantity
	_, err = service.productRepo.Update(product, lock)
	if err != nil {
		service.locker.Abort(lock)
		return nil, err
	}
	// update balance of user
	user.Balance -= trx.TotalPrice
	_, err = service.userRepo.Update(user, lock)
	if err != nil {
		service.locker.Abort(lock)
		return nil, err
	}
	// update sold amount of category
	category, err := service.categoryRepo.FindById(product.CategoryId, lock)
	if err != nil {
		service.locker.Abort(lock)
		return nil, err
	}
	category.SoldProductAmount += trx.Quantity
	_, err = service.categoryRepo.Update(category, lock)
	if err != nil {
		service.locker.Abort(lock)
		return nil, err
	}

	created, err := service.trxRepo.Create(trx, lock)
	if err != nil {
		service.locker.Abort(lock)
		return nil, err
	}

	service.locker.EndLock(lock)

	return created, nil
}
