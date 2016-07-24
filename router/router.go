package router

import (
	"github.com/labstack/echo"
	"jiminy_app/controller"
)

func Init() *echo.Echo {
	e := echo.New()

	// Routes
	v1 := e.Group("/jiminy/v1")
	{
		v1.GET("/history/list/:size", controller.GetHistory)
	}
	return e
}
