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
	products := &data.Products{}
	err := data.GetProducts(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = products.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

// AddProduct ti add new product to productsList
func AddProduct(w http.ResponseWriter, r *http.Request) {
	p := &data.Product{}
	if r.Body == nil {
		http.Error(w, "You must send data", http.StatusBadRequest)
		return
	}

	err := p.FromJSON(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validate the product
	err = p.Validate()
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Error validating product: %s", err),
			http.StatusBadRequest,
		)
		return
	}
	err = data.AddProduct(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p.ToJSON(w)
}

//GetProduct to retrieve a single product
func GetProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	p := &data.Product{}
	err := data.GetProduct(id, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.ToJSON(w)
}

//UpdateProduct handler
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	prod := &data.Product{}
	err := data.GetProduct(id, prod)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = prod.FromJSON(r.Body)

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
	// data.UpdateProduct(id, prod)

	err = data.UpdateProduct(prod)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Error when trying to update: %s", err),
			http.StatusBadRequest,
		)
		return
	}
	prod.ToJSON(w)

}

//DeleteProduct handler
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	prod := &data.Product{}

	if err := data.GetProduct(id, prod); err != nil {
		http.Error(
			w,
			fmt.Sprintf("%s", err),
			http.StatusBadRequest,
		)
		return
	}

	if err := data.DeleteProduct(prod); err != nil {
		http.Error(
			w,
			fmt.Sprintf("%s", err),
			http.StatusBadRequest,
		)
		return
	}
	_ = prod.ToJSON(w)
}
