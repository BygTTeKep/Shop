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
	p.HandleFunc("/", h.createProduct()).Methods("POST")
	u := router.PathPrefix("/user").Subrouter()
	u.HandleFunc("/{id:[0-9]+}/", h.deleteUser()).Methods("DELETE")
	//router.HandleFunc("/deleteUser/{id}", h.deleteUser()).Methods("DELETE")
	// router.HandleFunc("/createProduct", h.createProduct()).Methods("POST")
	// router.HandleFunc("/deleteProduct/{id}", h.deleteProduct()).Methods("DELETE")
	return router
}
