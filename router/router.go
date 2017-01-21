package router

import (
	"jiminy-app/controller"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	// Routes
	v1 := e.Group("/jiminy")
	{
		// Youtube動画情報
		v1.GET("/video/:id", controller.GetOneVideo)
		v1.GET("/video", controller.GetVideo)
		v1.POST("/video", controller.PostVideo)

		// 履歴情報API
		v1.GET("/history/:size", controller.GetHistory)
		v1.POST("/history", controller.PostDummy)
	}
	return e
}
