package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/TChi91/GoBuy/data"
	"github.com/go-chi/chi"
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

//GetProduct to retrieve a single product
func GetProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	prod, err := data.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	prod.ToJSON(w)
}

//UpdateProduct handler
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	prod := &data.Product{}
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
	data.UpdateProduct(id, prod)
	prod.ToJSON(w)

}

//DeleteProduct handler
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	products, err := data.DeleteProduct(id)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("%s", err),
			http.StatusBadRequest,
		)
		return
	}
	_ = products.ToJSON(w)
}
