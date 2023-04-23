package service

import "github.com/g91TeJl/Shop/pkg/repository"

type CartUser struct {
	repo repository.Cart
}

func NewcartUser(repo repository.Cart) *CartUser {
	return &CartUser{repo: repo}
}

//	func (c *CartUser) CreateCart(username, password string) (int, error) {
//		return c.repo.CreateCart(username, password)
//	}
func (c *CartUser) GetCart(id int) (int, error) {
	return c.repo.GetCart(id)
}
