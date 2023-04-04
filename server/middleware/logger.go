package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	//log实例化
	logger := log.New()
	return func(c *gin.Context) {
		//先占下位置
		fmt.Println(logger)
	}
}
