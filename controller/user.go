package controller

import (
	"net/http"
	"strconv"

	"go-orm-tutorial/model"

	"github.com/jinzhu/gorm"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

var _ CrudController = (*UserController)(nil)

// UserController endpoints
type UserController struct{}

// ReadAll endpoint
func (ctrl UserController) ReadAll(c echo.Context) error {
	return withDB(func(db *gorm.DB) error {
		users := make([]*model.User, 0)
		db.Find(&users)
		return c.JSONPretty(http.StatusOK, users, "  ")
	})
}

// Read endpoint
func (ctrl UserController) Read(c echo.Context) error {
	return withDB(func(db *gorm.DB) error {
		user := &model.User{}
		user.ID, _ = strconv.Atoi(c.Param("id"))
		db.First(user)
		return c.JSONPretty(http.StatusOK, user, "  ")
	})
}

// Create endpoint
func (ctrl UserController) Create(c echo.Context) (err error) {
	return withDB(func(db *gorm.DB) error {
		user := &model.User{}
		user.ID = 0
		if err = c.Bind(user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.(validator.ValidationErrors).Error())
		}
		db.Create(user)
		return c.JSONPretty(http.StatusOK, echo.Map{
			"message": "New User Successfully Created",
			"user":    user,
		}, " ")
	})
}

// Update endpoint
func (ctrl UserController) Update(c echo.Context) (err error) {
	return withDB(func(db *gorm.DB) error {
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
		return c.JSONPretty(http.StatusOK, echo.Map{
			"message": "User Successfully Update",
			"user":    user,
		}, " ")
	})
}

// Delete endpoint
func (ctrl UserController) Delete(c echo.Context) error {
	return withDB(func(db *gorm.DB) error {
		user := &model.User{}
		user.ID, _ = strconv.Atoi(c.Param("id"))
		db.First(user)
		db.Delete(user)
		return c.JSONPretty(http.StatusOK, echo.Map{
			"message": "User Successfully Delete",
			"user":    user,
		}, " ")
	})
}
