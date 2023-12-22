package data

import "tokbel/entity"

type TransactionRepoInterface interface {
	FindAll() []entity.TransactionHistory
	FindByUserId(userId int) []entity.TransactionHistory
	Create(trx *entity.TransactionHistory, lock *Lock) (*entity.TransactionHistory, error)
}
