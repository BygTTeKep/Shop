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

func TestHandlerProductCreate(t *testing.T) {
	db, teardown := TestingDB(t)
	defer teardown("products")
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handlers := handler.Newhandler(service)
	srv := new(Shop.Server)
	srv.Run(":8080", handlers.InitRoutes())
	defer srv.Shutdown(context.Background())
	product := model.TestProduct(t)
	testCase := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name:         "OK",
			payload:      product,
			expectedCode: http.StatusCreated,
		},
		{
			name:         "Bad",
			payload:      "",
			expectedCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/createProduct", b)
			handlers.InitRoutes().ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestHandlerDeleteProduct(t *testing.T) {
	db, teardown := TestingDB(t)
	defer teardown("products")
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handlers := handler.Newhandler(service)
	srv := new(Shop.Server)
	srv.Run(":8080", handlers.InitRoutes())
	defer srv.Shutdown(context.Background())
	product := model.TestProduct(t)
	id, _ := service.CreateProduct(*product)
	product.Id = id
	testCase := []struct {
		name         string
		payload      interface{} //поменять на тип model.Products или interface
		expectedCode int
	}{
		{
			name:         "OK",
			payload:      id,
			expectedCode: http.StatusOK,
		},
		{
			name:         "not Ok",
			payload:      123,
			expectedCode: http.StatusNotFound,
		},
		{
			name:         "not Ok",
			payload:      12,
			expectedCode: http.StatusNotFound,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			// b := &bytes.Buffer{}
			// json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/Product/%d/", tc.payload), nil)
			handlers.InitRoutes().ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)

		})
	}
}
