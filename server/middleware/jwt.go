package middleware

import (
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"net/http"
	"starblog/utils"
	"starblog/utils/errmsg"
	"strings"
)

// JWT 定义密钥参数
type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(utils.JwtKey),
	}
}

// MyClaims 定义结构体来接受参数
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken 生成token
func (j *JWT) GenerateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

// ParserToken 解析token
// todo 完善token错误处理
func (j *JWT) ParserToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.MalformedTokenError); ok {
			if err == ve {
				return nil, errmsg.TokenMalformed
			}
		}

	}

	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errmsg.TokenInvalid
	}
	return nil, errmsg.TokenInvalid
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		j := NewJWT()
		// 解析token
		claims, err := j.ParserToken(checkToken[1])
		if err != nil {
			if err == errmsg.TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status":  errmsg.ERROR,
					"message": "token授权已过期,请重新登录",
					"data":    nil,
				})
				c.Abort()
				return
			}
			// 其他错误
			c.JSON(http.StatusOK, gin.H{
				"status":  errmsg.ERROR,
				"message": err.Error(),
				"data":    nil,
			})
			c.Abort()
			return
		}

		c.Set("username", claims)

		c.Next()
	}
}
