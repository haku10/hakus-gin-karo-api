package main

import (
	"github.com/gin-gonic/gin"
    "net/http"
    "log"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
		"github.com/kelseyhightower/envconfig"
)

func main() {
    engine := gin.Default()
    engine.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "hello world",
        })
        })
        
    engine.GET("/test", func(c *gin.Context) {
        userName := regist()
        c.JSON(http.StatusOK, gin.H{
            "message": userName,
        })
        })

    engine.Run()
}

func regist() string {
	connect := gormConnect()
	defer connect.Close()
	user := User{}
	user.Name = "regist test"
    connect.Create(&user)
    return user.Name
}

// Env DBの環境情報
type Env struct {
	User     string      `env:"User" envDefault:"root"`
	Port     string      `env:"PORT" envDefault:"3306"`
	Pass     string      `env:"Pass" envDefault:"password"`
	Protocol string      `env:"Protocol" envDefault:"tcp(rest-db:3306)"`
	Dbname   string      `env:"Dbname" envDefault:"karo"`
}

func gormConnect() *gorm.DB {
    var dbinfo Env
	envconfig.Process("", &dbinfo)
	DBMS := "mysql"
	USER := dbinfo.User
	PASS := dbinfo.Pass
	PROTOCOL := dbinfo.Protocol
	DBNAME := dbinfo.Dbname

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
    log.Print(CONNECT)
	connection, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return connection
}

// User userテーブルの情報
type User struct {
	Id     int
	Name   string
	Age    int
	Gender int
}
