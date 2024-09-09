package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type router struct {
	router *mux.Router
}

// NewRouter :nodoc:
func NewRouter() Router {
	return &router{
		router: mux.NewRouter(),
	}
}

// Route :nodoc:
func (rtr *router) Route() *mux.Router {
	root := rtr.router.PathPrefix("/").Subrouter()
	root.HandleFunc("/health", healthCheckControler.Healthcek).Methods(http.MethodGet)
	in := root.PathPrefix("/in").Subrouter()
	productGroup := in.PathPrefix("/products").Subrouter()
	productGroup.HandleFunc("/items", itemController.ItemList).Methods(http.MethodGet)
	return rtr.router
}
