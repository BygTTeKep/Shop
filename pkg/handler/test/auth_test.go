package test_handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/g91TeJl/Shop"
	"github.com/g91TeJl/Shop/pkg/handler"
	"github.com/g91TeJl/Shop/pkg/model"
	"github.com/g91TeJl/Shop/pkg/repository"
	"github.com/g91TeJl/Shop/pkg/service"
	"github.com/stretchr/testify/assert"
)

func TestHandlerUserCreate(t *testing.T) {
	db, teardown := TestingDB(t)
	defer teardown("users")
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handlers := handler.Newhandler(service)
	srv := new(Shop.Server)
	srv.Run(":8080", handlers.InitRoutes())
	defer srv.Shutdown(context.Background())
	testCase := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "Valid",
			payload: map[string]string{
				"UserName":  "user",
				"FirstName": "firstName",
				"LastName":  "lastname",
				"Surname":   "surname",
				"Email":     "lastname@example.com",
				"Password":  "XXXXXXXXXXXXXX",
			},
			expectedCode: http.StatusCreated,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sign-up", b)
			handlers.InitRoutes().ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestHandlerUserLogin(t *testing.T) {
	db, teardown := TestingDB(t)
	defer teardown("users")
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handlers := handler.Newhandler(service)
	srv := new(Shop.Server)
	srv.Run(":8080", handlers.InitRoutes())
	defer srv.Shutdown(context.Background())
	u := model.TestUser(t)
	b := &bytes.Buffer{}
	json.NewEncoder(b).Encode(u)
	service.CreateUser(*u)
	id, err := service.GenerateToken(u.UserName, u.Password)
	if err != nil {
		t.Fatal(err)
		return
	}
	testCase := []struct {
		name         string
		cookieValue  map[string]interface{}
		expectedCode int
	}{
		{
			name: "authenticated",
			cookieValue: map[string]interface{}{
				"user_id": id,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "not authenticated",
			cookieValue:  nil,
			expectedCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/sign-in", b)
			req.Header.Set("Cookie", fmt.Sprintf("%s", tc.cookieValue["user_id"]))
			handlers.InitRoutes().ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestUserDelete(t *testing.T) {
	db, teardown := TestingDB(t)
	defer teardown("users")
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handlers := handler.Newhandler(service)
	srv := new(Shop.Server)
	srv.Run(":8080", handlers.InitRoutes())
	u := model.TestUser(t)
	id, _ := service.CreateUser(*u)
	testCase := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name:         "Delete",
			payload:      id,
			expectedCode: http.StatusOK,
		},
		{
			name:         "NotFound",
			payload:      -1,
			expectedCode: http.StatusNotFound,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/user/%d/delete/", tc.payload), nil)
			handlers.InitRoutes().ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestHandlerUserUpdate(t *testing.T) {
	db, teardown := TestingDB(t)
	defer teardown("users")
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handlers := handler.Newhandler(service)
	srv := new(Shop.Server)
	srv.Run(":8080", handlers.InitRoutes())
	u := model.TestUser(t)
	id, _ := service.CreateUser(*u)
	username, firstname, lastname, surname, email, password := "user", "firstname", "lastname", "surname", "example2@example.com", "XXXXXXXXXXXXXX"
	updateUser := model.UpdateUser{
		UserName:  &username,
		FirstName: &firstname,
		LastName:  &lastname,
		Surname:   &surname,
		Email:     &email,
		Password:  &password,
	}
	testCase := []struct {
		name        string
		payload     interface{}
		exeptedCode int
	}{
		{
			name:        "OK",
			payload:     updateUser,
			exeptedCode: http.StatusOK,
		},
		{
			name:        "BadRequest",
			payload:     model.UpdateUser{},
			exeptedCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/user/%d/update/", id), b)
			handlers.InitRoutes().ServeHTTP(rec, req)
			assert.Equal(t, tc.exeptedCode, rec.Code)
		})
	}
}

func TestHandlerGetAllProducts(t *testing.T) {
	db, teardown := TestingDB(t)
	defer teardown("products", "users", "cart", "cart_product")
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handlers := handler.Newhandler(service)
	srv := new(Shop.Server)
	srv.Run(":8080", handlers.InitRoutes())
	defer srv.Shutdown(context.Background())
	product := model.TestProduct(t)
	u := model.TestUser(t)
	idU, _ := service.CreateUser(*u)
	u.Id = idU
	id, _ := service.CreateProduct(*product)
	product.Id = id
	service.CreateCart(idU)
	service.AddProductToCart(idU, id)
	// product.Name = "asw"
	// id1, _ := service.CreateProduct(*product)
	// product.Id = id1
	// service.AddProductToCart(idU, id1)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/user/%d/cart/", idU), nil)
	handlers.InitRoutes().ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}
