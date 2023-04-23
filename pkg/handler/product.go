package handler

import (
	"encoding/json"
	"net/http"

	"github.com/g91TeJl/Shop/pkg/model"
)

func (h *Handler) createProduct() http.HandlerFunc {
	var req model.Products
	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			newErrorResponse(w, http.StatusBadRequest, err.Error())
		}
		if _, err := h.service.CreateProduct(req); err != nil {
			newErrorResponse(w, http.StatusInternalServerError, err.Error())
		}
		// if err := h.service.Products.AddProductPhoto()
	}
}

func (h *Handler) deleteProduct() http.HandlerFunc {
	var product model.Products
	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			id, err := h.service.GetProductId(product)
			if err != nil {
				newErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			if err := h.service.DeleteProduct(id); err != nil {
				newErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
		}
	}
}

func (h *Handler) AddProductPhoto() http.HandlerFunc {
	var req model.ProductPhoto
	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			newErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		if err := h.service.AddProductPhoto(req); err != nil {
			newErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
}
