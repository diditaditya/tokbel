package models

import (
	"gorm.io/gorm"
)

type TransactionHistory struct {
	gorm.Model
	ProductID  int
	UserID     int
	Quantity   int
	TotalPrice int
	Product    Product
	User       User
}
