package entity

import "time"

type Product struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryId int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	Category   Category  `json:"-"`
}
