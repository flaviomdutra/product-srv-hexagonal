package db

import (
	"database/sql"

	"github.com/flaviomdutra/product-srv-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("SELECT id, name, price, status FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.Id, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int
	p.db.QueryRow("SELECT COUNT(*) FROM products WHERE id = ?", product.GetId()).Scan(&rows)
	if rows == 0 {
		_, err := p.create(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	}

	return product, nil
}
func (p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare("INSERT INTO products VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {

	stmt, err := p.db.Prepare("UPDATE products SET name = ?, price = ?, status = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(product.GetName(), product.GetPrice(), product.GetStatus(), product.GetId())
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return product, nil
}
