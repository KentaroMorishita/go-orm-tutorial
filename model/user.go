package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	DB "go-orm-tutorial/db"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// User Model
type User struct {
	gorm.Model
	Name  string
	Email string
}

func init() {
	DB.InitialMigration(&User{})
}

// AllUsers endpoint
func (user User) AllUsers(w http.ResponseWriter, r *http.Request) {
	db := DB.ConnectDB()
	defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

// NewUser endpoint
func (user User) NewUser(w http.ResponseWriter, r *http.Request) {
	db := DB.ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "New User Successfully Created")
}

// DeleteUser endpoint
func (user User) DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := DB.ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	db.Where("name = $1", name).Debug().Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "User Successfully Delete")
}

// UpdateUser endpoint
func (user User) UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := DB.ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Where("name = $1", name).Debug().Find(&user)

	user.Email = email

	db.Save(&user)

	fmt.Fprintf(w, "User Successfully Update")
}
