package entity

import "time"

type Category struct {
	Id                int       `json:"id"`
	Type              string    `json:"type" validate:"required"`
	SoldProductAmount int       `json:"soldProductAmount"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Products          []Product `json:"products"`
}
