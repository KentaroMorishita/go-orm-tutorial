package controller

import (
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
