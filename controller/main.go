package controller

import (
	"github.com/labstack/echo"
)

// CrudController interface
type CrudController interface {
	Create(c echo.Context) (err error)
	ReadAll(c echo.Context) (err error)
	Read(c echo.Context) (err error)
	Update(c echo.Context) (err error)
	Delete(c echo.Context) (err error)
}
