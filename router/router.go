package router

import (
	"github.com/labstack/echo"
	"jiminy_app/controller"
)

func Init() *echo.Echo {
	e := echo.New()

	// Routes
	v1 := e.Group("/jiminy")
	{
		v1.GET("/history/:size", controller.GetHistory)
		v1.POST("/history", controller.PostDummy)
	}
	return e
}
