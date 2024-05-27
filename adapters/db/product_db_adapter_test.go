package db_test

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"

	"github.com/PedrodeAlmeidaFreitas/go-hex/adapters/db"
	"github.com/PedrodeAlmeidaFreitas/go-hex/application"
)

var Db *sql.DB
var tableName = "products"

func setup() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	query := fmt.Sprintf(`CREATE TABLE %s(
		"id" STRING, 
		"name" STRING,
		"price" FLOAT,
		"status" STRING
	);`, tableName)

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := fmt.Sprintf(`INSERT INTO %s 
	VALUES ("abc", "Product Test", 0, "disabled")
	`, tableName)
	stmt, err := db.Prepare(insert)

	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDbAdapter_Get(t *testing.T) {
	setup()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, float32(0.0), product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDbAdapter_Save(t *testing.T) {
	setup()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Enable()
	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, productResult.GetStatus(), application.ENABLED)
}
