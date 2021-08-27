package data

import (
	"encoding/json"
	"io"
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

type Products []*Product

//instead of putting the data extracting logic in handler,
//place the abstraction here instead

func GetProducts() Products {
	return productList
}

//method to convert the products to JSON format
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
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
