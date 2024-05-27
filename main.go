package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	dbadapter "github.com/PedrodeAlmeidaFreitas/go-hex/adapters/db"
	"github.com/PedrodeAlmeidaFreitas/go-hex/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "sqlite.db")
	productDb := dbadapter.NewProductDb(db)
	productService := application.NewProductService(productDb)

	product, err := productService.Create("Product Name", 30)
	if err != nil {
		fmt.Print("Error while adding new product ", err.Error())
	}
	productService.Enable(product)
}
