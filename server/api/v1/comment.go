package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"starblog/model"
	"starblog/service"
	"starblog/utils/errmsg"
	"strconv"
)

// AddComment 添加评论,引进结构体
func AddComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data)
	fmt.Println("1.", data)
	if data.Content != "" {
		code := service.CreateComment(&data)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		code := errmsg.ERROR_COMMENT_IS_EMPTY //评论不能为空
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// ShowComments 显示评论列表
func ShowComments(c *gin.Context) {
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
	//传给model中的ShowComments函数，返回一个comment切片
	data, totalNum := service.ShowComments(pageSize, pageNum)
	//将数据传递给前端展示
	c.JSON(http.StatusOK, gin.H{
		"status":   errmsg.SUCCESS,
		"TotalNum": totalNum,
		"data":     data,
		"message":  errmsg.GetErrMsg(errmsg.SUCCESS),
	})

}

// EditComment 编辑评论
func EditComment(c *gin.Context) {
	var data model.Comment
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := service.CheckCategory(data.Content)
	if code == errmsg.SUCCESS {
		code = service.EditComment(id, &data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// DelComment 删除评论
func DelComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := service.DelComment(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
