package v1

//导入gin包
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"starblog/model"
	"starblog/utils/errmsg"
)

var code int

// UserExist 查询用户是否存在
func UserExist(c *gin.Context) {

}

// AddUser 添加用户,引进结构体
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	} else if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	} else {
		code = errmsg.ERROR
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// ShowUsers 显示用户列表
func ShowUsers(c *gin.Context) {
	//pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	//pageNum, _ := strconv.Atoi(c.Query("pagenum"))

}

//查询单个用户

// EditUser 编辑用户
func EditUser(c *gin.Context) {

}

// DelUser 删除用户
func DelUser(c *gin.Context) {

}
