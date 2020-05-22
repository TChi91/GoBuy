package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TChi91/GoBuy/data"
)

//GetProducts return productsList
func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	result, _ := json.Marshal(products)
	fmt.Fprint(w, string(result))
}

// AddProduct ti add new product to productsList
func AddProduct(w http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}

	err := json.NewDecoder(r.Body).Decode(prod)

	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}
	data.AddProduct(prod)
}
