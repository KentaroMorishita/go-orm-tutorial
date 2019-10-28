package controller

import (
	"net/http"
	"strconv"
	"time"

	"go-orm-tutorial/env"
	"go-orm-tutorial/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

var _ CrudController = (*UserController)(nil)

// UserController endpoints
type UserController struct{}

func generateJwtToken(user *model.User) (string, error) {
	jwtKey := env.Get("JWT_SECRET", "")
	jwtLifeTime := time.Hour * 72

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.JwtClaims{
		user.ID,
		user.Name,
		user.Email,
		user.IsAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtLifeTime).Unix(),
		},
	})
	return token.SignedString([]byte(jwtKey))
}

// Login endpoint
func (ctrl UserController) Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	return withDB(func(db *gorm.DB) error {
		user := &model.User{}
		db.Find(&user, "email=?", email)

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			return echo.ErrUnauthorized
		}

		token, err := generateJwtToken(user)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{"token": token})
	})

}

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
