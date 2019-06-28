package model

import (
	DB "go-orm-tutorial/db"
	"time"
)

// User Model
type User struct {
	ID        int        `json:"id" gorm:"primary_key"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

func init() {
	DB.InitialMigration(&User{})
}
