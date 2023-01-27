package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"starblog/utils"
)

// InitRouter 创建公用类
func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default() //此处可以用Default或者New，区别是Default默认加了两个中间件-日志文件和错误恢复
	//初始路由，由于采用前后端分离以及版本控制，所以使用路由组
	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}
	err := r.Run(utils.HttpPort)
	if err != nil {
		fmt.Println("路由初始化失败，请检查:", err)
	}
}
