package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/g91TeJl/Shop/pkg/model"
	"github.com/gorilla/mux"
)

func (h *Handler) signUp() http.HandlerFunc {
	var req model.User
	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			newErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		// json.NewEncoder(w).Encode(id)
		// if id, err := h.service.Authorization.CreateUser(req); err != nil {
		// 	newErrorResponse(w, http.StatusInternalServerError, err.Error())
		// 	return
		// }
		id, err := h.service.Authorization.CreateUser(req)
		if err != nil {
			newErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		if _, err := h.service.Cart.CreateCart(id); err != nil {
			newErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		/*
			+ метод для добавления в продукта в корзину и подсчет итогового чека
		*/
		w.WriteHeader(http.StatusCreated)
	}
}
func (h *Handler) signIn() http.HandlerFunc {
	type signInInput struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var req signInInput
	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			newErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		token, err := h.service.Authorization.GenerateToken(req.Username, req.Password)
		if err != nil {
			newErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		r.Header.Add("Cookie", token)

		json.NewEncoder(w).Encode(token)
	}
}

func (h *Handler) deleteUser() http.HandlerFunc { //переписать в принятие id а не asfasf
	return func(w http.ResponseWriter, r *http.Request) {
		// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// 	newErrorResponse(w, http.StatusBadRequest, err.Error())
		// 	return
		// }
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			newErrorResponse(w, http.StatusBadRequest, err.Error())
		}
		err = h.service.Authorization.DeleteUser(id)
		if err != nil {
			newErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
