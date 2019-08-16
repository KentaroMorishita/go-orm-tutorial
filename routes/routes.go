package routes

import (
	"fmt"
	ctrl "go-orm-tutorial/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
)

// Router is setup routes
func Router() {
	e := initEcho()
	api := e.Group("/api")
	{
		crud(api, &ctrl.UserController{}, "users")
	}

	e.Logger.Fatal(e.Start(":8000"))
}

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func initEcho() (e *echo.Echo) {
	e = echo.New()
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-CSRF-TOKEN",
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = &customValidator{validator: validator.New()}
	return e
}

func crud(api *echo.Group, c ctrl.CrudController, prefix string) {
	g := api.Group(fmt.Sprintf("/%s", prefix))
	g.POST("", c.Create)
	g.GET("", c.ReadAll)
	g.GET("/:id", c.Read)
	g.PUT("/:id", c.Update)
	g.DELETE("/:id", c.Delete)
}
