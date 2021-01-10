package main
import (
    "github.com/gin-gonic/gin"
    "net/http"
    "math/rand"
    "strconv"
    _"github.com/jinzhu/gorm/dialects/mysql"
	"ginrest/app/service"
    "ginrest/app/model"
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
    var num = strconv.Itoa(rand.Intn(100))
    user.Name = "regist test" + num
    connect.Create(&user)
    return user.Name
}
