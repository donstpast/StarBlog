package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"starblog/middleware"
	"starblog/model"
	"starblog/service"
	"starblog/utils/errmsg"
	"strconv"
)

// AddArticle 添加文章,引进结构体
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR_TOKEN_NOT_EXIST,
			"massage": errmsg.GetErrMsg(errmsg.ERROR_TOKEN_NOT_EXIST),
		})
		c.Abort()
		return
	}
	name := username.(*middleware.MyClaims).Username
	id, code := service.ShowUserID(name)
	if data.Uid == 0 {
		data.Uid = int(id)
	}

	code = service.CreateArticle(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// todo ShowArticle 显示文章列表

func ShowArticles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	title := c.Query("title")
	//如果pageSize或者pageNum为0，则进行gorm中的 Cancel limit condition with -1
	//例子：db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = 1
	}
	//传给model中的ShowArticles函数，返回一个user切片
	data, code, totalNum := service.ShowArticles(title, pageSize, pageNum)
	//将数据传递给前端展示
	c.JSON(http.StatusOK, gin.H{
		"status":   code,
		"TotalNum": totalNum,
		"data":     data,
		"message":  errmsg.GetErrMsg(code),
	})

}

// ShowCategoryArticles 查询单个分类下所有的文章
func ShowCategoryArticles(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
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
	//传给model中的ShowCategoryArticles函数，返回一个user切片
	data, code, totalNum := service.ShowCategoryArticles(id, pageSize, pageNum)
	//将数据传递给前端展示
	c.JSON(http.StatusOK, gin.H{
		"status":   code,
		"TotalNum": totalNum,
		"data":     data,
		"message":  errmsg.GetErrMsg(code),
	})
}

// ShowSingleArticle 显示单个文章内容
func ShowSingleArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := service.ShowSingleArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

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
