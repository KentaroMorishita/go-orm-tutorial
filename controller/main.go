package controller

import (
	DB "go-orm-tutorial/db"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// CrudController interface
type CrudController interface {
	Create(c echo.Context) error
	ReadAll(c echo.Context) error
	Read(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

func withDB(fn func(db *gorm.DB) error) error {
	db := DB.ConnectDB()
	defer db.Close()
	return fn(db)
}
