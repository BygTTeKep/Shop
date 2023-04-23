package repository_test

import (
	"testing"

	"github.com/g91TeJl/Shop/pkg/model"
	"github.com/g91TeJl/Shop/pkg/repository"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryCreate(t *testing.T) {
	db, teardown := TestingDB(t)
	defer teardown("users")
	repo := repository.NewRepository(db)
	u := model.TestUser(t)
	_, err := repo.CreateUser(*u)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	_, err = repo.GetUser(u.UserName, u.Password)
	assert.NoError(t, err)
}

func TestUserRepositoryDelete(t *testing.T) {
	db, teardown := TestingDB(t)
	defer teardown("users")
	repo := repository.NewRepository(db)
	u := model.TestUser(t)
	_, err := repo.CreateUser(*u)
	assert.NoError(t, err)
	err = repo.DeleteUser(u.UserName, u.Password)
	assert.NoError(t, err)
}
