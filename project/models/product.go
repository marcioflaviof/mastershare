package models

type Product struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}

type UpdatableProduct struct {
	Filter SecureProducts `json: "filter"`
	Update Product        `json: "update"`
}

type SecureProducts struct {
	Name string `json:"name" validate:"required"`
}
