package service

import (
	"github.com/g91TeJl/Shop/pkg/model"
	"github.com/g91TeJl/Shop/pkg/repository"
)

type CartUser struct {
	repo repository.Cart
}

func NewcartUser(repo repository.Cart) *CartUser {
	return &CartUser{repo: repo}
}

func (c *CartUser) CreateCart(idU int) (int, error) {
	return c.repo.CreateCart(idU)
}
func (c *CartUser) GetCart(id int) (int, error) {
	return c.repo.GetCart(id)
}
func (c *CartUser) AddProductToCart(id int, idProduct int) (int, error) {
	return c.repo.AddProductToCart(id, idProduct)
}

func (c *CartUser) GetAllProductFromCartProducts(cart_id int) ([]model.Products, error) {
	return c.repo.GetAllProductFromCartProducts(cart_id)
}
