package repository

import (
	"fmt"

	"github.com/g91TeJl/Shop/pkg/model"
	"github.com/jmoiron/sqlx"
)

type CartUser struct {
	db      *sqlx.DB
	auth    Authorization
	product Product
}

func NewCart(db *sqlx.DB) *CartUser {
	return &CartUser{db: db}
}

func (c *CartUser) CreateCart(idU int) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s(customer_id) values ($1) RETURNING id", cartTable)
	row := c.db.QueryRow(query, idU)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (c *CartUser) GetCart(id int) (int, error) {
	var idCart int
	query := fmt.Sprintf("SELECT id FROM %s WHERE customer_id=$1", cartTable)
	row := c.db.QueryRow(query, id)
	if err := row.Scan(&idCart); err != nil {
		return -1, err
	}
	return idCart, nil
}

func (c *CartUser) AddProductToCart(id int, idProduct int) (int, error) {
	cart_id, err := c.GetCart(id)
	if err != nil {
		return -1, err
	}
	// product_id, err := c.product.GetProductId(product)
	// if err != nil {
	// 	return -1, err
	// }
	query := fmt.Sprintf("INSERT INTO %s(cart_id, product_id) values($1,$2)", cartProductTable)
	//row := c.db.QueryRow(query, cart_id, idProduct)
	_, err = c.db.Exec(query, cart_id, idProduct)
	return cart_id, err
	/*
		Добавлнение в таблицу cart_product
		get cart_id, product_id
	*/
}
func (c *CartUser) GetAllProductFromCartProducts(cart_id int) ([]model.Products, error) {
	var Values int64
	setValues := make([]int64, 0)
	//var allProduct []model.Products
	//var Values2 string
	query := fmt.Sprintf("SELECT product_id FROM %s WHERE cart_id=$1", cartProductTable)
	row, err := c.db.Query(query, cart_id)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		if err := row.Scan(&Values); err != nil {
			return nil, err
		}
		setValues = append(setValues, Values)
	}
	allProduct := make([]model.Products, len(setValues))
	//setQuery := strings.Join(setValues, ",")
	//queryP := fmt.Sprintf("SELECT name, description, price FROM %s WHERE id IN (%s)", productsTable, setQuery)
	for i, id := range setValues {
		queryp := fmt.Sprintf("SELECT name,description, price FROM %s WHERE id=$1", productsTable)
		// if err := c.db.Get(&allProduct, queryp, id); err != nil {
		// 	return nil, err
		// }
		c.db.QueryRow(queryp, id).Scan(&allProduct[i].Name, &allProduct[i].Description, &allProduct[i].Price)
		// queryp = fmt.Sprintf("SELECT description FROM %s WHERE id=$1", productsTable)
		// if err := c.db.Get(allProduct[i].Description, queryp, id); err != nil {
		// 	return nil, err
		// }
		// queryp = fmt.Sprintf("SELECT price FROM %s WHERE id=$1", productsTable)
		// if err := c.db.Get(allProduct[i].Price, queryp, id); err != nil {
		// 	return nil, err
		// }
	}
	// queryP := fmt.Sprintf("SELECT name, description, price FROM %s WHERE id=$%s", productsTable, setQuery)
	// if err := c.db.Get(&allProduct, queryP); err != nil {
	// 	return nil, err
	// }
	return allProduct, nil
}
