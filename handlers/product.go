package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/TChi91/GoBuy/data"
	"github.com/TChi91/GoBuy/db"
	"github.com/go-chi/chi"
)

//GetProducts return productsList
func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := &data.Products{}
	if err := db.Db.Find(&products).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := products.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	// fmt.Fprint(w, string(err))
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
	if err := db.Db.Create(p).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p.ToJSON(w)
}

//GetProduct to retrieve a single product
func GetProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	p := &data.Product{}

	if err := db.Db.First(&p, "id = ?", id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.ToJSON(w)
}

//UpdateProduct handler
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	prod := &data.Product{}
	if err := db.Db.First(&prod, "id = ?", id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	// data.UpdateProduct(id, prod)
	db.Db.Save(prod)
	db.Db.Updates(prod)
	prod.ToJSON(w)

}

//DeleteProduct handler
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	prod := &data.Product{}
	if err := db.Db.First(&prod, "id = ?", id).Error; err != nil {
		http.Error(
			w,
			fmt.Sprintf("%s", err),
			http.StatusBadRequest,
		)
		return
	}
	db.Db.Delete(prod)

	_ = prod.ToJSON(w)
}
