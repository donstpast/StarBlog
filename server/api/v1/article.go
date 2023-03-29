package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"starblog/model"
	"starblog/service"
	"starblog/utils/errmsg"
	"strconv"
)

//显示文章列表

//查询单个文章

//编辑文章

//删除文章

// AddArticle 添加文章,引进结构体
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code := service.CreateArticle(&data)
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// todo ShowArticle 显示文章列表

func ShowArticles(c *gin.Context) {
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
	data := service.ShowArticles(pageSize, pageNum)
	//将数据传递给前端展示
	c.JSON(http.StatusOK, gin.H{
		"status":  errmsg.SUCCESS,
		"data":    data,
		"message": errmsg.GetErrMsg(errmsg.SUCCESS),
	})

}

//todo 查询单个分类下所有的文章

//todo 显示单个文章内容

// EditArticle 编辑文章
func EditArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := service.EditArticle(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// DelArticle 删除分类
func DelArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := service.DelArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
