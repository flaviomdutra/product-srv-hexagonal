package main

import (
	"database/sql"

	db2 "github.com/flaviomdutra/product-srv-hexagonal/adapters/db"
	"github.com/flaviomdutra/product-srv-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "sqlite.db")

	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Product 2", 30.0)
	productService.Enable(product)
}
