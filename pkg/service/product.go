package service

import (
	"github.com/g91TeJl/Shop/pkg/model"
	"github.com/g91TeJl/Shop/pkg/repository"
)

type Product struct {
	repo repository.Products
}

func NewProducts(repo repository.Products) *Product {
	return &Product{repo: repo}
}

func (s *Product) CreateProduct(product model.Products) (int, error) {
	return s.repo.CreateProduct(product)
}
func (s *Product) GetProductId(product model.Products) (int, error) {
	return s.repo.GetProductId(product)
}

func (s *Product) DeleteProduct(id int) error {
	return s.repo.DeleteProduct(id)
}

func (s *Product) AddProductPhoto(productPhoto model.ProductPhoto) error {
	return s.repo.AddProductPhoto(productPhoto)
}
func (s *Product) GetProductById(id int) error {
	return s.repo.GetProductById(id)
}

func (s *Product) UpdateProductInput(id int, input model.UpdateProduct) error {
	return s.repo.UpdateProductInput(id, input)
}
