package product

type ProductRequest struct {
	Title      string `json:"title" validate:"required"`
	Price      int    `json:"price" validate:"required,gte=0,lte=50000000"`
	Stock      int    `json:"stock" validate:"required,gte=5"`
	CategoryId int    `json:"category_id" validate:"required"`
}
