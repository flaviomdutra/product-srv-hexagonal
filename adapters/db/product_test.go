package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/flaviomdutra/product-srv-hexagonal/adapters/db"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `
		CREATE TABLE products (
		 	"id" string,
	    "name" string,
			"price" float,
			"status" string
		);
	`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products VALUES ("1", "Product 1", 10.0, "disabled");`

	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDbGet(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product, err := productDb.Get("1")

	require.Nil(t, err)
	require.Equal(t, "1", product.GetId())
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 10.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}
