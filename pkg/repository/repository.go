package repository

import (
	"github.com/g91TeJl/Shop/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
	UpdateUser(id int, user model.User) error
	DeleteUser(id int) error
}

type Products interface {
	CreateProduct(product model.Products) (int, error)
	DeleteProduct(id int) error
	GetProductId(product model.Products) (int, error)
	AddProductPhoto(product_photo model.ProductPhoto) error
	GetProductById(id int) error
	UpdateProductInput(id int, input model.UpdateProduct) error
}

type Cart interface {
	CreateCart(idU int) (int, error)
	AddProductToCart(id int, idProduct int) (int, error)
	GetCart(id int) (int, error)
	GetAllProductFromCartProducts(cart_id int) ([]model.Products, error)
}

type Repository struct {
	Authorization
	Products
	Cart
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Products:      NewProduct(db),
		Cart:          NewCart(db),
	}
}
