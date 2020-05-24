package data

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-playground/validator"
)

//ErrNotFound => product not found
var ErrNotFound = fmt.Errorf("Product not found")

// Product structure
type Product struct {
	ID          int     `json:"id"`
	Brand       string  `json:"brand" validate:"required"`
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"desc" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

//Products type used in Get Products
type Products []*Product

//Validate product
func (p *Product) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

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
func (p Products) ToJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	err := enc.Encode(p)
	return err
}

//GetProduct func
func GetProduct(id int) (*Product, error) {
	prod, _, err := findProduct(id)
	if err != nil {
		return nil, ErrNotFound
	}
	return prod, nil
}

//UpdateProduct func
func UpdateProduct(id int, p *Product) error {
	_, pos, _ := findProduct(id)
	p.ID = id
	productList[pos] = p
	return nil
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrNotFound
}

/*
// func findProduct(id int) (*Product, int, error) {
// 	for i := 1; i <= len(productList); i++ {
// 		fmt.Println(i, id, productList[i].ID)
// 		if productList[i-1].ID == id {
// 			return productList[i], i, nil
// 		}
// 	}
// 	return nil, -1, ErrNotFound
// }
*/

//DeleteProduct func
func DeleteProduct(id int) (Products, error) {
	_, pos, _ := findProduct(id)
	if pos == -1 {
		return nil, ErrNotFound
	}
	productList = append(productList[:pos], productList[pos+1:]...)
	for _, p := range productList {
		out := fmt.Sprintf("ID %v\n", p.ID)
		fmt.Println(out)
	}
	fmt.Println(len(productList))
	return productList, nil
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

var productList = []*Product{
	{
		ID:          1,
		Brand:       "Fujitsu",
		Title:       "Fujitsu Laptop",
		Description: "The best laptop ever",
		Price:       999.00,
	},
	{
		ID:          2,
		Brand:       "Canon",
		Title:       "Canon Printer",
		Description: "The best printer ever",
		Price:       100.00,
	},
}
