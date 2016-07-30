package main

import (
	"github.com/labstack/echo/engine/fasthttp"
	"jiminy_app/router"
)

func main() {
	route := router.Init()
	route.Run(fasthttp.New(":61970"))
}