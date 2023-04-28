package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/g91TeJl/Shop/pkg/model"
	"github.com/gorilla/mux"
)

func (h *Handler) createProduct() http.HandlerFunc {
	var req model.Products
	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			newErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		req.Validate()
		if _, err := h.service.CreateProduct(req); err != nil {
			newErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		// if err := h.service.Products.AddProductPhoto()
		w.WriteHeader(http.StatusCreated)
	}
}

func (h *Handler) deleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		// 	newErrorResponse(w, http.StatusBadRequest, err.Error())
		// 	return
		// }
		// id, err := h.service.GetProductId(product)
		// if err != nil {
		// 	newErrorResponse(w, http.StatusBadRequest, err.Error())
		// 	return
		// }
		// vars := mux.Vars(r)
		// id, _ := strconv.Atoi(vars["id"])
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			newErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		err = h.service.GetProductById(id)
		if err != nil {
			newErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		if err := h.service.DeleteProduct(id); err != nil {
			newErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
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

func (h *Handler) updateProduct() http.HandlerFunc {
	var product model.UpdateProduct
	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			newErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			newErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		err = product.Validate()
		if err != nil {
			newErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		err = h.service.UpdateProductInput(id, product)
		if err != nil {
			newErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
