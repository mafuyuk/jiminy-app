package main

import (
	"jiminy-app/router"
)

func main() {
	route := router.Init()
	route.Start(":61970")
}