package transaction

import "tokbel/entity"

type CreateTransactionResponse struct {
	Message         string          `json:"message"`
	TransactionBill TransactionBill `json:"transaction_bill"`
}

type TransactionBill struct {
	TotalPrice   int    `json:"total_price"`
	Quantity     int    `json:"quantity"`
	ProductTitle string `json:"product_title"`
}

type UserTransaction struct {
	Id         int            `json:"id"`
	ProductId  int            `json:"product_id"`
	UserId     int            `json:"user_id"`
	Quantity   int            `json:"quantity"`
	TotalPrice int            `json:"total_price"`
	Product    entity.Product `json:"product"`
}
