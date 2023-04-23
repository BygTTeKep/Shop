package repository_test

import (
	"fmt"
	"testing"

	"github.com/g91TeJl/Shop/pkg/model"
	"github.com/g91TeJl/Shop/pkg/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreateCart(t *testing.T) {
	db, teardown := TestingDB(t)
	defer teardown("cart", "users")
	repo := repository.NewRepository(db)
	u := *model.TestUser(t)
	id, err := repo.CreateUser(u)
	fmt.Println("succes")
	_, err = repo.CreateCart(id)
	assert.NoError(t, err)
}

func TestAddProductToCart(t *testing.T) {
	db, teardown := TestingDB(t)
	defer teardown("cart_product", "users", "products")
	repo := repository.NewRepository(db)
	u := *model.TestUser(t)
	p := model.TestProduct(t)
	id, err := repo.CreateUser(u)
	assert.NoError(t, err)
	_, err = repo.CreateCart(id)
	assert.NoError(t, err)
	idProduct, err := repo.CreateProduct(*p)
	assert.NoError(t, err)
	_, err = repo.AddProductToCart(id, idProduct)
	assert.NoError(t, err)
}
