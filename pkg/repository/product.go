package repository

import (
	"fmt"
	"strings"

	"github.com/g91TeJl/Shop/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Product struct {
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

func (r *Product) GetProductId(product model.Products) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE name=$1 AND description=$2 AND price=$3", productsTable)
	row := r.db.QueryRow(query, product.Name, product.Description, product.Price)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

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

func (r *Product) UpdateProductInput(id int, input model.Products) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, input.Name)
		argId++
	}
	if input.Description != "" {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, input.Description)
		argId++
	}
	if input.Price >= 1 {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, input.Price)
		argId++
	}
	setQuery := strings.Join(setValues, ",")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", productsTable, setQuery, argId)
	args = append(args, id)
	_, err := r.db.Exec(query, args...)
	return err
}
