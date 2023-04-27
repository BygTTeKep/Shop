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
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return idCart, nil
}

/*
	Просмотр продуктов в корзине (работа с
	увеличении кол-во товаров	(cart_product
*/

func (c *CartUser) AddProductToCart(id int, idProduct int) (int, error) {
	cart_id, err := c.GetCart(id)
	if err != nil {
		return -1, err
	}
	// product_id, err := c.product.GetProductId(product)
	// if err != nil {
	// 	return -1, err
	// }
	var idPr int
	query := fmt.Sprintf("INSERT INTO %s(cart_id, product_id) values($1,$2)", cartProductTable)
	row := c.db.QueryRow(query, cart_id, idProduct)
	if err := row.Scan(&id); err != nil {
		return -1, nil
	}
	return idPr, nil
	/*
		Добавлнение в таблицу cart_product
		get cart_id, product_id
	*/
}
func (c *CartUser) GetAllProductFromCart(cartId int) ([]model.Products, error) {
	var allProduct []model.Products
	query := fmt.Sprintf("SELECT * FROM %s WHERE cart_id=$1", cartProductTable) //Запрос неверный
	if err := c.db.Get(&allProduct, query, cartId); err != nil {
		return nil, err
	}
	return allProduct, nil
}
