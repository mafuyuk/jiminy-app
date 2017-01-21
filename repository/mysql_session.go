package repository

import (
	"fmt"

	"jiminy-app/util"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

const sourceFormat string = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true"

func InitMysql() error {
	conf, err := util.GetDBConfig()
	if err != nil {
		return err
	}
	fmt.Println(conf)

	dataSource := fmt.Sprintf(sourceFormat, conf.User, conf.Pass, conf.Host, conf.Port, "test")
	db, err = gorm.Open("mysql", dataSource)
	if err != nil {
		return err
	}
	return nil
}

func CloseMysql() {
	db.Close()
}
