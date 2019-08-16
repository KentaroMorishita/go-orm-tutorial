package routes

import (
	"fmt"
	ctrl "go-orm-tutorial/controller"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

// CustomValidator is validator
type CustomValidator struct {
	validator *validator.Validate
}

// Validate is custom validation
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// Router create router
func Router() (e *echo.Echo) {
	e = echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	api := e.Group("/api")
	crud(api, &ctrl.UserController{}, "users")
	return
}

func crud(api *echo.Group, c ctrl.CrudController, prefix string) {
	g := api.Group(fmt.Sprintf("/%s", prefix))
	g.POST("", c.Create)
	g.GET("", c.ReadAll)
	g.GET("/:id", c.Read)
	g.PUT("/:id", c.Update)
	g.DELETE("/:id", c.Delete)
}
