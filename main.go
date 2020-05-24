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
	r.Get("/{id}", handlers.GetProduct)
	r.Put("/{id}", handlers.UpdateProduct)
	r.Delete("/{id}", handlers.DeleteProduct)
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()
}
