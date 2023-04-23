package handler

import (
	"encoding/json"
	"net/http"

	"github.com/g91TeJl/Shop/pkg/model"
)

func (h *Handler) signUp() http.HandlerFunc {
	var req model.User
	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			newErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		id, err := h.service.Authorization.CreateUser(req)
		if err != nil {
			newErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		json.NewEncoder(w).Encode(id)
		// if id, ids, err := h.service.Authorization.CreateUser(req); err != nil {
		// 	newErrorResponse(w, http.StatusInternalServerError, err.Error())
		// 	return
		// }
		// if _, err := h.service.Cart.CreateCart(req.UserName, req.Password); err != nil {
		// 	newErrorResponse(w, http.StatusInternalServerError, err.Error())
		// 	return //В CreateUser Добавить CreateCart
		// }
		/*
			Собственный обработчик ошибок
			+ respond
			= создать корзину при создании пользователя
			+ метод для добавления в продукта в корзину и подсчет итогового чека
		*/
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
		json.NewEncoder(w).Encode(token)
	}
}

func (h *Handler) deleteUser() http.HandlerFunc {
	type deleteUserInput struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var req deleteUserInput
	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			newErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		err := h.service.Authorization.DeleteUser(req.Username, req.Password)
		if err != nil {
			newErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
}
