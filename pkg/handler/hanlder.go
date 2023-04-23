package handler

import (
	"github.com/g91TeJl/Shop/pkg/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
}

func Newhandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.Use(h.commonMiddleware)
	router.HandleFunc("/sign-up", h.signUp()).Methods("POST")
	router.HandleFunc("/sign-in", h.signIn()).Methods("POST")
	router.HandleFunc("/deleteUser", h.deleteUser()).Methods("DELETE")
	router.HandleFunc("/createProduct", h.createProduct()).Methods("POST") //название
	router.HandleFunc("/deleteProduct", h.deleteProduct()).Methods("DELETE")
	return router
}
