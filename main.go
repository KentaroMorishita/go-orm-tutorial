package main

import (
	"fmt"
	"log"
	"net/http"

	"go-orm-tutorial/model"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func handleRequests() {

	user := &model.User{}

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/users", user.AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", user.NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}/{email}", user.UpdateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{name}", user.DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("Go ORM Tutorial")
	handleRequests()
}
