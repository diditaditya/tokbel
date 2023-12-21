package category

import "time"

type CategoryResponse struct {
	Id                int       `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount int       `json:"sold_product_amount"`
	CreatedAt         time.Time `json:"created_at"`
}
