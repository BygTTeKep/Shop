package repository_test

import (
	"testing"

	"github.com/g91TeJl/Shop/pkg/model"
	"github.com/g91TeJl/Shop/pkg/repository"
	"github.com/stretchr/testify/assert"
)

func TestProductCreate(t *testing.T) {
	db, teardown := TestingDB(t)
	defer teardown("products")
	product := model.TestProduct(t)
	repo := repository.NewRepository(db)
	id, err := repo.CreateProduct(*product)
	assert.NoError(t, err)
	assert.NotNil(t, id)
}

func TestProductGetId(t *testing.T) {
	db, teardown := TestingDB(t)
	defer teardown("products")
	product := model.TestProduct(t)
	repo := repository.NewRepository(db)
	_, err := repo.CreateProduct(*product)
	assert.NoError(t, err)
	id, err := repo.GetProductId(*product)
	assert.NoError(t, err)
	assert.NotNil(t, id)
}

func TestProductDelete(t *testing.T) {
	db, teardown := TestingDB(t)
	defer teardown("products")
	product := model.TestProduct(t)
	repo := repository.NewRepository(db)
	id, err := repo.CreateProduct(*product)
	assert.NoError(t, err)
	assert.NotNil(t, id)
	err = repo.DeleteProduct(id)
	assert.NoError(t, err)
}

func TestProductUpdate(t *testing.T) {
	db, teardown := TestingDB(t)
	defer teardown("products")
	product := model.TestProduct(t)
	repo := repository.NewRepository(db)
	id, err := repo.CreateProduct(*product)
	assert.NoError(t, err)
	assert.NotNil(t, id)
	name := "test"
	productUpdate := model.UpdateProduct{}
	productUpdate.Name = &name
	err = repo.UpdateProductInput(id, productUpdate)
	assert.NoError(t, err)
}
