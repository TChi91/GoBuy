package data

import (
	"encoding/json"
	"io"
)

// Product structure
type Product struct {
	ID          int
	Brand       string
	Title       string
	Description string
	Price       float64
}

//Products type used in Get Products
type Products []*Product

// GetProducts returns all products from the database
func GetProducts() Products {
	return productList
}

//AddProduct append new product to productsList
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

//ToJSON for marshaling productsList
func (p *Products) ToJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	err := enc.Encode(p)
	return err
}

//ToJSON for marshaling product
func (p *Product) ToJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	err := enc.Encode(p)
	return err
}

//FromJSON for marshaling product
func (p *Product) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	err := dec.Decode(p)
	return err
}

func getNextID() int {
	lastID := len(productList)
	nextID := lastID + 1
	return nextID
}

var productList = Products{
	&Product{
		ID:          1,
		Brand:       "Fujitsu",
		Title:       "Fujitsu Laptop",
		Description: "The best laptop ever",
		Price:       999.00,
	},
	&Product{
		ID:          2,
		Brand:       "Canon",
		Title:       "Canon Printer",
		Description: "The best printer ever",
		Price:       100.00,
	},
}
