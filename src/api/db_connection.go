// package main

// import (
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// )

// // Env DBの環境情報
// type Env struct {
// 	User     string      `env:"User" envDefault:"root"`
// 	Port     string      `env:"PORT" envDefault:"3306"`
// 	Pass     string      `env:"Pass" envDefault:"password"`
// 	Protocol string      `env:"Protocol" envDefault:"tcp(rest-db:3306)"`
// 	Dbname   string      `env:"Dbname" envDefault:"karo"`
// }

// func gormConnect() *gorm.DB {
// 	dbinfo := Env{}
// 	DBMS := "mysql"
// 	USER := dbinfo.User
// 	PASS := dbinfo.Pass
// 	PROTOCOL := dbinfo.Protocol
// 	DBNAME := dbinfo.Dbname

// 	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
// 	connection, err := gorm.Open(DBMS, CONNECT)

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return connection
// }
