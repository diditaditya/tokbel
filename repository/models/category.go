package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Type              string
	SoldProductAmount int
	Products          []Product
}
