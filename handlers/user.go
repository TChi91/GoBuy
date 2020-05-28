package handlers

import (
	"net/http"

	"github.com/TChi91/GoBuy/data"
)

//AddUser to add new product to productsList
func AddUser(w http.ResponseWriter, r *http.Request) {
	user := &data.User{}
	if r.Body == nil {
		http.Error(w, "You must send data", http.StatusBadRequest)
		return
	}

	err := user.FromJSON(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validate the product
	// err = p.Validate()
	// if err != nil {
	// 	http.Error(
	// 		w,
	// 		fmt.Sprintf("Error validating product: %s", err),
	// 		http.StatusBadRequest,
	// 	)
	// 	return
	// }
	user = user.Clean()
	err = data.Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.ToJSON(w)
}

// //GetProduct to retrieve a single product
// func GetProduct(w http.ResponseWriter, r *http.Request) {
// 	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
// 	p := &data.Product{}
// 	err := data.GetProduct(id, p)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	p.ToJSON(w)
// }

// //UpdateProduct handler
// func UpdateProduct(w http.ResponseWriter, r *http.Request) {
// 	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
// 	prod := &data.Product{}
// 	err := data.GetProduct(id, prod)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	err = prod.FromJSON(r.Body)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// validate the product
// 	err = prod.Validate()
// 	if err != nil {
// 		http.Error(
// 			w,
// 			fmt.Sprintf("Error validating product: %s", err),
// 			http.StatusBadRequest,
// 		)
// 		return
// 	}
// 	// data.UpdateProduct(id, prod)

// 	err = data.UpdateProduct(prod)
// 	if err != nil {
// 		http.Error(
// 			w,
// 			fmt.Sprintf("Error when trying to update: %s", err),
// 			http.StatusBadRequest,
// 		)
// 		return
// 	}
// 	prod.ToJSON(w)

// }

// DeleteProduct handler
// func DeleteProduct(w http.ResponseWriter, r *http.Request) {
// 	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
// 	prod := &data.Product{}

// 	if err := data.GetProduct(id, prod); err != nil {
// 		http.Error(
// 			w,
// 			fmt.Sprintf("%s", err),
// 			http.StatusBadRequest,
// 		)
// 		return
// 	}

// 	if err := data.DeleteProduct(prod); err != nil {
// 		http.Error(
// 			w,
// 			fmt.Sprintf("%s", err),
// 			http.StatusBadRequest,
// 		)
// 		return
// 	}
// 	_ = prod.ToJSON(w)
// }
