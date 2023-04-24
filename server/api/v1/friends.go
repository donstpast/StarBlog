package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"starblog/model"
	"starblog/service"
	"starblog/utils/errmsg"
	"strconv"
)

// AddFriends 添加友链,引进结构体
func AddFriends(c *gin.Context) {
	var data model.Friends
	_ = c.ShouldBindJSON(&data)
	if data.Link != "" {
		code := service.CreateFriends(&data)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		code := errmsg.ERROR_FRIENDS_LINKS_IS_EMPTY //友链为空
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// ShowFriends 显示友链列表
func ShowFriends(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	//如果pageSize或者pageNum为0，则进行gorm中的 Cancel limit condition with -1
	//例子：db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = 1
	}
	//传给model中的ShowFriends函数，返回一个friend切片
	data, totalNum := service.ShowFriends(pageSize, pageNum)
	//将数据传递给前端展示
	c.JSON(http.StatusOK, gin.H{
		"status":   errmsg.SUCCESS,
		"TotalNum": totalNum,
		"data":     data,
		"message":  errmsg.GetErrMsg(errmsg.SUCCESS),
	})

}

// EditFriends 编辑友链
func EditFriends(c *gin.Context) {
	var data model.Friends
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := service.EditFriends(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// DelFriends 删除友链
func DelFriends(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := service.DelFriends(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
