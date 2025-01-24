package dto

type CreateProductDTO struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	InStock     bool    `json:"in_stock"`
}
