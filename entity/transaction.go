package entity

import "time"

type TransactionHistory struct {
	Id         int       `json:"id"`
	ProductId  int       `json:"product_id"`
	UserId     int       `json:"user_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice int       `json:"total_price"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	Product    Product   `json:"product"`
	User       User      `json:"-"`
}
