package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "starblog/api/v1"
	"starblog/utils"
)

// InitRouter 创建公用类
func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default() //此处可以用Default或者New，区别是Default默认加了两个中间件-日志文件和错误恢复
	//初始路由，由于采用前后端分离以及版本控制，所以使用路由组
	routerV1 := r.Group("api/v1")
	{
		//User模块的路由接口
		//添加用户
		routerV1.POST("user/add", v1.AddUser)
		//查看用户列表
		routerV1.GET("users", v1.ShowUsers)
		//编辑用户
		routerV1.PUT("user/:id", v1.EditUser)
		//删除用户
		routerV1.DELETE("user/:id", v1.DelUser)

		//Article模块的路由接口
		//添加文章
		routerV1.POST("article/add", v1.AddArticle)
		//查看文章列表
		routerV1.GET("articles", v1.ShowArticles)
		//查看单个文章内容
		routerV1.GET("article/:id", v1.ShowSingleArticle)
		//编辑文章
		routerV1.PUT("article/:id", v1.EditArticle)
		//删除文章
		routerV1.DELETE("article/:id", v1.DelArticle)

		//Category模块的路由接口
		//添加分类
		routerV1.POST("category/add", v1.AddCategory)
		//查看分类列表
		routerV1.GET("categories", v1.ShowCategories)
		//查看单个分类下文章
		routerV1.GET("category/articles/:id", v1.ShowCategoryArticles)
		//编辑分类
		routerV1.PUT("category/:id", v1.EditCategory)
		//删除分类
		routerV1.DELETE("category/:id", v1.DelCategory)

	}
	err := r.Run(utils.HttpPort)
	if err != nil {
		fmt.Println("路由初始化失败，请检查:", err)
	}
}
