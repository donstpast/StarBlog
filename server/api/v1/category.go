package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"starblog/model"
	"starblog/service"
	"starblog/utils/errmsg"
	"strconv"
)

// AddCategory 添加分类,引进结构体

func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	if data.Name != "" {
		code := service.CheckCategory(data.Name)
		if code == errmsg.SUCCESS {
			code = service.CreateCategory(&data)
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		code := errmsg.ERROR_CATEGORY_IS_EMPTY
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	}

}

// ShowCategories 显示分类列表
func ShowCategories(c *gin.Context) {
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
	data := service.ShowCategories(pageSize, pageNum)
	//将数据传递给前端展示
	c.JSON(http.StatusOK, gin.H{
		"status":  errmsg.SUCCESS,
		"data":    data,
		"message": errmsg.GetErrMsg(errmsg.SUCCESS),
	})

}

//todo 查询单个分类下所有的文章

// EditCategory 编辑分类
func EditCategory(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := service.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		code = service.EditCategory(id, &data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// DelCategory 删除分类
func DelCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := service.DelCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
