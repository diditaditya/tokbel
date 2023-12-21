package entity

import "time"

type TransactionHistory struct {
	Id         int
	ProductId  int
	UserId     int
	Quantity   int
	TotalPrice int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Product    Product
	User       User
}
