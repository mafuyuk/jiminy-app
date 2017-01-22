package main

import (
	"jiminy-app/repository"
	"jiminy-app/router"
	"log"
)

func main() {
	// MySQLコネクション生成
	if err := repository.InitMysql(); err != nil {
		log.Println("failed start mysql session: ", err)
		return
	}
	defer repository.CloseMysql()

	route := router.Init()
	route.Start(":61970")
}
