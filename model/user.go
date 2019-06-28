package model

import (
	DB "go-orm-tutorial/db"
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