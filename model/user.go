package model

import (
	DB "go-orm-tutorial/db"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JwtClaims jwt custom claims
type JwtClaims struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
	jwt.StandardClaims
}

// User Model
type User struct {
	ID        int        `json:"id" gorm:"primary_key"`
	Name      string     `json:"name" validate:"required"`
	Email     string     `json:"email" validate:"required,email"`
	Password  string     `json:"-"`
	IsAdmin   bool       `json:"is_admin"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

func init() {
	DB.InitialMigration(&User{})
}
