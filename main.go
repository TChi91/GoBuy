package main

import (
	"net/http"

	"github.com/TChi91/GoBuy/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.GetProducts)
	mux.HandleFunc("/new", handlers.AddProduct)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
