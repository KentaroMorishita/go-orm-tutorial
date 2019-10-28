package routes

import (
	"fmt"
	ctrl "go-orm-tutorial/controller"

	"go-orm-tutorial/env"
	"go-orm-tutorial/model"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
)

var jwtConfig = middleware.JWTConfig{
	Claims:     &model.JwtClaims{},
	SigningKey: []byte(env.Get("JWT_SECRET", "")),
}

// Router is setup routes
func Router() {
	e := initEcho()
	api := e.Group("/api")

	// login
	api.POST("/login", (&ctrl.UserController{}).Login)

	// users endpoints
	users := api.Group(fmt.Sprintf("/%s", "users"))
	users.Use(middleware.JWTWithConfig(jwtConfig))
	crud(users, &ctrl.UserController{})

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
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = &customValidator{validator: validator.New()}
	return e
}

func crud(g *echo.Group, c ctrl.CrudController) {
	g.POST("", c.Create)
	g.GET("", c.ReadAll)
	g.GET("/:id", c.Read)
	g.PUT("/:id", c.Update)
	g.DELETE("/:id", c.Delete)
}
