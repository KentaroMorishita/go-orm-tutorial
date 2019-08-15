package controller

import (
	"net/http"
	"strconv"

	DB "go-orm-tutorial/db"
	"go-orm-tutorial/model"

	"github.com/labstack/echo"
)

var _ CrudController = (*UserController)(nil)

// UserController endpoints
type UserController struct{}

// ReadAll endpoint
func (ctrl UserController) ReadAll(c echo.Context) (err error) {
	db := DB.ConnectDB()
	defer db.Close()

	users := make([]*model.User, 0)
	db.Find(&users)
	return c.JSON(http.StatusOK, users)
}

// Read endpoint
func (ctrl UserController) Read(c echo.Context) (err error) {
	db := DB.ConnectDB()
	defer db.Close()

	user := &model.User{}
	user.ID, _ = strconv.Atoi(c.Param("id"))
	db.First(user)
	return c.JSON(http.StatusOK, user)
}

// Create endpoint
func (ctrl UserController) Create(c echo.Context) (err error) {
	db := DB.ConnectDB()
	defer db.Close()

	user := &model.User{}
	user.ID = 0
	if err = c.Bind(user); err != nil {
		return
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
		return
	}
	db.Save(user)
	return c.String(http.StatusOK, "User Successfully Update")
}

// Delete endpoint
func (ctrl UserController) Delete(c echo.Context) (err error) {
	db := DB.ConnectDB()
	defer db.Close()

	user := &model.User{}
	user.ID, _ = strconv.Atoi(c.Param("id"))
	db.First(user)
	db.Delete(user)
	return c.String(http.StatusOK, "User Successfully Delete")
}
