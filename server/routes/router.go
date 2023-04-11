package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "starblog/api/v1"
	"starblog/middleware"
	"starblog/utils"
)

// InitRouter 创建公用类
func InitRouter() {
	gin.SetMode(utils.AppMode)
	//计划采用logrus保存为日志文件，同时使用gin自带的日志在终端显示
	r := gin.Default() //此处可以用Default或者New，区别是Default默认加了两个中间件-日志文件和错误恢复
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	//初始路由，由于采用前后端分离以及版本控制，所以使用路由组
	//分为权限组路由，和公共路由
	authV1 := r.Group("api/v1")
	authV1.Use(middleware.JwtToken())
	{
		//文件上传
		authV1.POST("/upload", v1.FileUploadLocal)

		//User模块的路由接口
		user := authV1.Group("/user")
		{
			//查看用户列表
			user.GET("/list", v1.ShowUsers)
			//编辑用户
			user.PUT("/:id", v1.EditUser)
			//删除用户
			user.DELETE("/:id", v1.DelUser)
		}
		//Article模块的路由接口
		article := authV1.Group("/article")
		{
			//添加文章
			article.POST("/add", v1.AddArticle)
			//编辑文章
			article.PUT("/:id", v1.EditArticle)
			//删除文章
			article.DELETE("/:id", v1.DelArticle)
		}
		//Category模块的路由接口
		category := authV1.Group("/category")
		{
			//添加分类
			category.POST("/add", v1.AddCategory)
			//编辑分类
			category.PUT("/:id", v1.EditCategory)
			//删除分类
			category.DELETE("/:id", v1.DelCategory)
		}

	}
	//公开路由组
	pubilcV1 := r.Group("api/v1")
	{
		//管理员登录接口
		pubilcV1.POST("/login", v1.AdminLogin)
		//User模块的路由接口
		//添加用户
		pubilcV1.POST("/register", v1.AddUser)
		//Article模块的路由接口
		//查看文章列表
		pubilcV1.GET("articles", v1.ShowArticles)
		//查看单个文章内容
		pubilcV1.GET("article/:id", v1.ShowSingleArticle)
		//Category模块的路由接口

		//查看分类列表
		pubilcV1.GET("categories", v1.ShowCategories)
		//查看单个分类下文章
		pubilcV1.GET("category/articles/:id", v1.ShowCategoryArticles)

	}
	err := r.Run(utils.HttpPort)
	if err != nil {
		fmt.Println("路由初始化失败，请检查:", err)
	}
}
