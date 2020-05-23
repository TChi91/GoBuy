package handlers

import (
	"fmt"
	"net/http"

	"github.com/TChi91/GoBuy/data"
)

//GetProducts return productsList
func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := products.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	// fmt.Fprint(w, string(err))
}

// AddProduct ti add new product to productsList
func AddProduct(w http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	if r.Body == nil {
		http.Error(w, "You must send data", http.StatusBadRequest)
		return
	}

	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validate the product
	err = prod.Validate()
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Error validating product: %s", err),
			http.StatusBadRequest,
		)
		return
	}
	data.AddProduct(prod)
	prod.ToJSON(w)
}
