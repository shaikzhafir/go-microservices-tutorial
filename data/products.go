package data

import (
	"time"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	SKU         string
	CreatedOn   string
	UpdatedOn   string
	DeletedOn   string
}

//instead of putting the data extracting logic in handler,
//place the abstraction here instead

func GetProducts() []*Product {
	return productList
}

//an array with a pointer
var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Milk Coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          1,
		Name:        "Espresso",
		Description: "Just Coffee",
		Price:       1.11,
		SKU:         "abc456",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
