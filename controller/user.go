package controller

import (
	"net/http"
	"strconv"

	DB "go-orm-tutorial/db"
	"go-orm-tutorial/model"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

var _ CrudController = (*UserController)(nil)

// UserController endpoints
type UserController struct{}

// ReadAll endpoint
func (ctrl UserController) ReadAll(c echo.Context) error {
	db := DB.ConnectDB()
	defer db.Close()

	users := make([]*model.User, 0)
	db.Find(&users)
	return c.JSONPretty(http.StatusOK, users, "  ")
}

// Read endpoint
func (ctrl UserController) Read(c echo.Context) error {
	db := DB.ConnectDB()
	defer db.Close()

	user := &model.User{}
	user.ID, _ = strconv.Atoi(c.Param("id"))
	db.First(user)
	return c.JSONPretty(http.StatusOK, user, "  ")
}

// Create endpoint
func (ctrl UserController) Create(c echo.Context) (err error) {
	db := DB.ConnectDB()
	defer db.Close()

	user := &model.User{}
	user.ID = 0
	if err = c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.(validator.ValidationErrors).Error())
	}
	db.Create(user)
	return c.String(http.StatusOK, "New User Successfully Created")
}

// Update endpoint
func (ctrl UserController) Update(c echo.Context) (err error) {
	db := DB.ConnectDB()
	defer db.Close()

	user := &model.User{}
	user.ID, _ = strconv.Atoi(c.Param("id"))
	db.First(user)
	if err = c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.(validator.ValidationErrors).Error())
	}
	db.Save(user)
	return c.String(http.StatusOK, "User Successfully Update")
}

// Delete endpoint
func (ctrl UserController) Delete(c echo.Context) error {
	db := DB.ConnectDB()
	defer db.Close()

	user := &model.User{}
	user.ID, _ = strconv.Atoi(c.Param("id"))
	db.First(user)
	db.Delete(user)
	return c.String(http.StatusOK, "User Successfully Delete")
}
