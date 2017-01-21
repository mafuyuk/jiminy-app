package main

import (
	"jiminy-app/repository"
	"jiminy-app/router"
)

func main() {
	// MySQLコネクション生成
	repository.InitMysql()
	defer repository.CloseMysql()

	route := router.Init()
	route.Start(":61970")
}
