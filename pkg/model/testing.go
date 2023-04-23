package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		UserName:  "user",
		FirstName: "firstName",
		LastName:  "lastname",
		Surname:   "surname",
		Email:     "lastname@example.com",
		Password:  "XXXXXXXXXXXXXX",
	}
}

func TestProduct(t *testing.T) *Products {
	return &Products{
		Name:        "name",
		Description: "description",
		Price:       100.0,
	}
}

func TestProductPhoto(t *testing.T) *ProductPhoto {
	return &ProductPhoto{
		Url:        "Url",
		Product_id: 1,
	}
}
