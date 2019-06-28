package routes

import (
	"go-orm-tutorial/controller"

	"github.com/gorilla/mux"
)

// Router create router
func Router() *mux.Router {
	userController := &controller.UserController{}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", userController.AllUsers).Methods("GET")
	router.HandleFunc("/user/{name}/{email}", userController.NewUser).Methods("POST")
	router.HandleFunc("/user/{name}/{email}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{name}", userController.DeleteUser).Methods("DELETE")

	return router
}
