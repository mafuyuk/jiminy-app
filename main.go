package main

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Run(standard.New(":1323"))
}
