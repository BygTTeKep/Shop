package service

import (
	"github.com/g91TeJl/Shop/pkg/model"
	"github.com/g91TeJl/Shop/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	DeleteUser(username, password string) error
}

type Products interface {
	CreateProduct(product model.Products) (int, error)
	DeleteProduct(id int) error
	GetProductId(product model.Products) (int, error)
	AddProductPhoto(product_photo model.ProductPhoto) error
}

type Cart interface {
	//CreateCart(username, password string) (int, error) //перенести в CreateUser
	GetCart(id int) (int, error)
}

type Service struct {
	Authorization
	Products
	Cart
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Products:      NewProducts(repo.Products),
		Cart:          NewcartUser(repo.Cart),
	}
}
