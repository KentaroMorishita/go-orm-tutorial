package main

import (
	"fmt"
	"log"
	"net/http"

	"go-orm-tutorial/controller"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func handleRequests() {

	userController := &controller.UserController{}

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/users", userController.AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", userController.NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}/{email}", userController.UpdateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{name}", userController.DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("Go ORM Tutorial")
	handleRequests()
}
