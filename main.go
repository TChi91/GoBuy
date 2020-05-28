package main

import (
	"net/http"

	"github.com/TChi91/GoBuy/data"
	"github.com/TChi91/GoBuy/db"
	"github.com/TChi91/GoBuy/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	err := db.Open()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Db.AutoMigrate(&data.Product{})
	db.Db.AutoMigrate(&data.User{})

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handlers.GetProducts)
	r.Post("/", handlers.AddProduct)
	r.Get("/{id}", handlers.GetProduct)
	r.Put("/{id}", handlers.UpdateProduct)
	r.Delete("/{id}", handlers.DeleteProduct)

	r.Post("/register", handlers.AddUser)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()
}
