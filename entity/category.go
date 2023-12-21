package entity

import "time"

type Category struct {
	Id                int       `json:"id"`
	Type              string    `json:"type" validate:"required"`
	SoldProductAmount int       `json:"soldProductAmount"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	Products          []Product
}
