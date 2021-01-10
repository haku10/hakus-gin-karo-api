package main
import (
    "github.com/gin-gonic/gin"
    "net/http"
    _"github.com/jinzhu/gorm/dialects/mysql"
	"ginrest/app/service"
    "ginrest/app/model"
    "time"
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
    connect := service.GormConnect()
    defer connect.Close()
    user := model.User{}
    user.Name = "regist test"
    connect.Create(&user)
    return user.Name
}
