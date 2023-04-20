package v1

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"net/http"
	"starblog/middleware"
	"starblog/model"
	"starblog/service"
	"starblog/utils/errmsg"
	"time"
)

// AdminLogin 后台登录
func AdminLogin(c *gin.Context) {
	var data model.User
	var code int
	_ = c.ShouldBindJSON(&data)
	data, code = service.CheckLogin(data.Username, data.Password)
	fmt.Println(code)
	if code == errmsg.SUCCESS {
		token := setToken(c, data)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data.Username,
			"id":      data.ID,
			"message": errmsg.GetErrMsg(code),
			"token":   token,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data.Username,
			"id":      data.ID,
			"message": errmsg.GetErrMsg(code),
		})
	}

}

// token生成函数
func setToken(c *gin.Context, user model.User) string {
	j := middleware.NewJWT()

	claims := middleware.MyClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: jwt.NewTime(float64(time.Now().Unix() - 100)),
			ExpiresAt: jwt.NewTime(float64(time.Now().Unix() + 86400)),
			ID:        string(user.ID),
			Issuer:    "StarBlog",
		},
	}
	token, err := j.GenerateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR,
			"message": errmsg.GetErrMsg(errmsg.ERROR),
			"token":   token,
		})
	}
	return token
}
