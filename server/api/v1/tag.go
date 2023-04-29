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

// AddTag 添加标签,引进结构体
func AddTag(c *gin.Context) {
	var data model.Tag
	_ = c.ShouldBindJSON(&data)
	if data.Name != "" {
		code := service.CheckTag(data.Name)
		if code == errmsg.SUCCESS {
			code = service.CreateTag(&data)
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		code := errmsg.ERROR_TAG_IS_EMPTY
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	}

}

// ShowTag 显示标签列表
func ShowTag(c *gin.Context) {
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
	//传给model中的ShowUser函数，返回一个user切片
	data, totalNum := service.ShowTags(pageSize, pageNum)
	//将数据传递给前端展示
	c.JSON(http.StatusOK, gin.H{
		"status":   errmsg.SUCCESS,
		"TotalNum": totalNum,
		"data":     data,
		"message":  errmsg.GetErrMsg(errmsg.SUCCESS),
	})

}

// ShowTagInfo 查询单个标签
func ShowTagInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := service.ShowTagInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditTag 编辑标签
func EditTag(c *gin.Context) {
	var data model.Tag
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := service.CheckTag(data.Name)
	if code == errmsg.SUCCESS {
		code = service.EditTag(id, &data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// DelTag 删除标签
func DelTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := service.DelTag(id)
	fmt.Println("1.", code)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
