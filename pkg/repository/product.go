package repository

import (
	"fmt"

	"github.com/g91TeJl/Shop/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Product struct { //сменить название
	db *sqlx.DB
}

func NewProduct(db *sqlx.DB) *Product {
	return &Product{db: db}
}

func (r *Product) CreateProduct(product model.Products) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s(name, description, price) values($1,$2,$3) RETURNING id", productsTable)
	row := r.db.QueryRow(query, product.Name, product.Description, product.Price)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (r *Product) GetProductId(product model.Products) (int, error) { //GetProductId
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE name=$1, description=$2, price=$3", productsTable)
	row := r.db.QueryRow(query, product.Name, product.Description, product.Price)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

//Добавить GetAllProduct для seller'a

func (r *Product) DeleteProduct(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", productsTable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Product) AddProductPhoto(productPhoto model.ProductPhoto) error {
	query := fmt.Sprintf("INSERT INTO %s(url, product_id) values($1,$2)", productPhotoTable)
	_, err := r.db.Exec(query, productPhoto.Url, productPhoto.Product_id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Product) GetProductById(id int) error {
	var product model.Products
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", productsTable)
	err := r.db.Get(&product, query, id)
	return err
}
