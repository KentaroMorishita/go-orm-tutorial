package routes

import (
	"go-orm-tutorial/controller"

	"github.com/gorilla/mux"
)

// Router create router
func Router() (router *mux.Router) {
	router = mux.NewRouter().StrictSlash(true)
	crud(router, &controller.UserController{}, "users")
	return
}

func crud(router *mux.Router, ctrl controller.CrudController, prefix string) {
	r := router.PathPrefix("/users").Subrouter()
	r.HandleFunc("", ctrl.Create).Methods("POST")
	r.HandleFunc("", ctrl.ReadAll).Methods("GET")
	r.HandleFunc("/{id}", ctrl.Read).Methods("GET")
	r.HandleFunc("/{id}", ctrl.Update).Methods("PUT")
	r.HandleFunc("/{id}", ctrl.Delete).Methods("DELETE")
}
