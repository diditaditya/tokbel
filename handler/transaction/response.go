package transaction

type CreateTransactionResponse struct {
	Message         string          `json:"message"`
	TransactionBill TransactionBill `json:"transaction_bill"`
}

type TransactionBill struct {
	TotalPrice   int    `json:"total_price"`
	Quantity     int    `json:"quantity"`
	ProductTitle string `json:"product_title"`
}
