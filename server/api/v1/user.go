package v1

//导入gin包
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"starblog/model"
	"starblog/service"
	"starblog/utils/errmsg"
	"starblog/utils/validator"
	"strconv"
)

// AddUser 添加用户,引进结构体
// 增加后端判断用户名和密码是否为空，若为空则不允许创建

func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	msg, validCode := validator.Validate(&data)
	if validCode != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  validCode,
			"massage": msg,
		})
		c.Abort()
		return
	}
	//if data.Username != "" && data.Password != "" {
	//	code := service.CheckUser(data.Username)
	//	if code == errmsg.SUCCESS {
	//		service.CreateUser(&data)
	//	} else if code == errmsg.ERROR_USERNAME_USED {
	//		code = errmsg.ERROR_USERNAME_USED
	//	} else {
	//		code = errmsg.ERROR
	//	}
	code := service.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		service.CreateUser(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//else {
//	code := errmsg.ERROR_USERNAME_OR_PASSWORD_IS_EMPTY
//	c.JSON(http.StatusBadRequest, gin.H{
//		"status":  code,
//		"message": errmsg.GetErrMsg(code),
//	})

//}

// ShowUsers 显示用户列表
func ShowUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	userName := c.Query("username")
	//如果pageSize或者pageNum为0，则进行gorm中的 Cancel limit condition with -1
	//例子：db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = 1
	}
	//传给model中的ShowUser函数，返回一个user切片
	data, totalNum := service.ShowUsers(userName, pageSize, pageNum)
	//将数据传递给前端展示
	c.JSON(http.StatusOK, gin.H{
		"status":   errmsg.SUCCESS,
		"TotalNum": totalNum,
		"data":     data,
		"message":  errmsg.GetErrMsg(errmsg.SUCCESS),
	})

}

//查询单个用户

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := service.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		code = service.EditUser(id, &data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// DelUser 删除用户
func DelUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := service.DelUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
