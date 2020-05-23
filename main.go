package main

import (
	"net/http"

	"github.com/TChi91/GoBuy/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handlers.GetProducts)
	r.Post("/", handlers.AddProduct)
	r.Put("/{id}", handlers.UpdateProduct)
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()
}
