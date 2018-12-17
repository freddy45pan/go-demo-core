package product

import (
	"math/rand"
)

// Product struct
type Product struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
}

// Read
func (product *Product) Read() (err error) {
	product.Title = "prod_" + product.ID
	product.Price = rand.Intn(100)
	return
}
