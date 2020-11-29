package service

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func GormConnect() *gorm.DB {
  DBMS := "mysql"
  USER := "root"
  PASS := "password"
  PROTOCOL := "tcp(rest-db:3306)"
  DBNAME := "karo"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	log.Print(CONNECT)
	connection, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return connection
}
