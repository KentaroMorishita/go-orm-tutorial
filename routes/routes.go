package routes

import (
	"fmt"
	ctrl "go-orm-tutorial/controller"

	"github.com/labstack/echo"
)

// Router create router
func Router() (e *echo.Echo) {
	e = echo.New()
	crud(e, &ctrl.UserController{}, "users")
	return
}

func crud(e *echo.Echo, c ctrl.CrudController, prefix string) {
	r := e.Group(fmt.Sprintf("/%s", prefix))
	r.POST("", c.Create)
	r.GET("", c.ReadAll)
	r.GET("/:id", c.Read)
	r.PUT("/:id", c.Update)
	r.DELETE("/:id", c.Delete)
}
