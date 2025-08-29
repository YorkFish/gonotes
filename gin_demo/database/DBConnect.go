package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConnect *gorm.DB

var err error

func ConnMysql() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gintest?charset=utf8mb4&parseTime=True&loc=Local"
	DBConnect, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
