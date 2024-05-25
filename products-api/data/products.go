package data

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	SKU         string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Product 1",
		Description: "Desc 1",
		Price:       12.99,
		SKU:         "demo",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		DeletedAt:   "",
	},
	&Product{
		ID:          2,
		Name:        "Product 2",
		Description: "Desc 2",
		Price:       12.2,
		SKU:         "demo",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		DeletedAt:   "",
	},
}
