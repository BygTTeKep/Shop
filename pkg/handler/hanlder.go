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
	p := router.PathPrefix("/Product").Subrouter()
	p.HandleFunc("/{id:[0-9]+}/", h.deleteProduct()).Methods("DELETE")
	//p.HandleFunc("/{id:[0-9]+}/", h.getProduct()).Methods("GET")
	p.HandleFunc("/{id:[0-9]+}/update/", h.updateProduct()).Methods("PUT")
	p.HandleFunc("/", h.createProduct()).Methods("POST")
	u := router.PathPrefix("/user").Subrouter()
	u.HandleFunc("/{id:[0-9]+}/delete/", h.deleteUser()).Methods("DELETE")
	u.HandleFunc("/{id:[0-9]+}/update/", h.updateUser()).Methods("PUT")
	u.HandleFunc("/{id:[0-9]+}/cart/", h.getProduct()).Methods("GET")
	u.HandleFunc("/{id:[0-9]+}/addProduct/{idProduct:[0-9]+}", h.addProduct()).Methods("POST")
	//router.HandleFunc("/deleteUser/{id}", h.deleteUser()).Methods("DELETE")
	// router.HandleFunc("/createProduct", h.createProduct()).Methods("POST")
	// router.HandleFunc("/deleteProduct/{id}", h.deleteProduct()).Methods("DELETE")
	return router
}
