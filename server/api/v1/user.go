package v1

//导入gin包
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"starblog/model"
	"starblog/utils/errmsg"
	"strconv"
)

var code int

// AddUser 添加用户,引进结构体
// 增加后端判断用户名和密码是否为空，若为空则不允许创建

func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	if data.Username != "" && data.Password != "" {
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
	} else {
		code = errmsg.ERROR_USERNAME_OR_PASSWORD_IS_EMPTY
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	}

}

// ShowUsers 显示用户列表
func ShowUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	//如果pageSize或者pageNum为0，则进行gorm中的 Cancel limit condition with -1
	//例子：db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	//传给model中的ShowUser函数，返回一个user切片
	data := model.ShowUsers(pageSize, pageNum)
	//将数据传递给前端展示
	c.JSON(http.StatusOK, gin.H{
		"status":  errmsg.SUCCESS,
		"data":    data,
		"message": errmsg.GetErrMsg(errmsg.SUCCESS),
	})

}

//查询单个用户

// EditUser 编辑用户
func EditUser(c *gin.Context) {

}

// DelUser 删除用户
func DelUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DelUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
