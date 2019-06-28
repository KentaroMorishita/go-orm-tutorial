package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	DB "go-orm-tutorial/db"
	"go-orm-tutorial/model"

	"github.com/gorilla/mux"
)

// UserController endpoints
type UserController struct{}

// AllUsers endpoint
func (ctrl UserController) AllUsers(w http.ResponseWriter, r *http.Request) {
	db := DB.ConnectDB()
	defer db.Close()

	var users []model.User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

// NewUser endpoint
func (ctrl UserController) NewUser(w http.ResponseWriter, r *http.Request) {
	db := DB.ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	name, email := vars["name"], vars["email"]

	db.Create(&model.User{Name: name, Email: email})
	fmt.Fprintf(w, "New User Successfully Created")
}

// DeleteUser endpoint
func (ctrl UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := DB.ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user model.User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)
	fmt.Fprintf(w, "User Successfully Delete")
}

// UpdateUser endpoint
func (ctrl UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := DB.ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	name, email := vars["name"], vars["email"]

	var user model.User
	db.Where("name = ?", name).Find(&user)
	user.Email = email
	db.Save(&user)
	fmt.Fprintf(w, "User Successfully Update")
}
