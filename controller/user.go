package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	DB "go-orm-tutorial/db"
	"go-orm-tutorial/model"

	"github.com/gorilla/mux"
)

// UserController endpoints
type UserController struct{}

// ReadAll endpoint
func (ctrl UserController) ReadAll(w http.ResponseWriter, r *http.Request) {
	db := DB.ConnectDB()
	defer db.Close()

	var users []model.User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

// Read endpoint
func (ctrl UserController) Read(w http.ResponseWriter, r *http.Request) {
	db := DB.ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	user := &model.User{}
	user.ID, _ = strconv.Atoi(vars["id"])
	db.First(&user)
	json.NewEncoder(w).Encode(user)
}

// Create endpoint
func (ctrl UserController) Create(w http.ResponseWriter, r *http.Request) {
	db := DB.ConnectDB()
	defer db.Close()

	user := &model.User{}
	user.ID = 0
	_ = json.NewDecoder(r.Body).Decode(&user)
	db.Create(user)
	fmt.Fprintf(w, "New User Successfully Created")
}

// Update endpoint
func (ctrl UserController) Update(w http.ResponseWriter, r *http.Request) {
	db := DB.ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)

	user := &model.User{}
	user.ID, _ = strconv.Atoi(vars["id"])
	db.First(&user)
	_ = json.NewDecoder(r.Body).Decode(&user)
	db.Save(&user)
	fmt.Fprintf(w, "User Successfully Update")
}

// Delete endpoint
func (ctrl UserController) Delete(w http.ResponseWriter, r *http.Request) {
	db := DB.ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)

	user := &model.User{}
	user.ID, _ = strconv.Atoi(vars["id"])
	db.First(&user)
	db.Delete(&user)
	fmt.Fprintf(w, "User Successfully Delete")
}
