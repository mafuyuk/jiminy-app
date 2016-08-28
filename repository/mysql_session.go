package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

var	db *gorm.DB
const sourceFormat string = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true"

func InitMysql() error {
	var err error
	dataSource := fmt.Sprintf(sourceFormat, "root", "1111", "192.168.99.100", "3306", "test")
	db, err = gorm.Open("mysql", dataSource)
	if err != nil {
		return err
	}
	return nil
}

func CloseMysql() {
	db.Close()
}