package entity

import "time"

type User struct {
	Id        int       `json:"id"`
	FullName  string    `json:"full_name" validate:"required"`
	Email     string    `json:"email" validate:"required, email"`
	Password  string    `json:"-"`
	Role      string    `json:"role" validate:"required, oneof='admin customer'"`
	Balance   int       `json:"balance" validate:"required, gte=0, lte=100000000"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
