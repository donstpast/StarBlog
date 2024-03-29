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
			//查看用户信息
			user.GET("/:id", v1.ShowUserInfo)
			//编辑用户
			user.PUT("/:id", v1.EditUser)
			//删除用户
			user.DELETE("/:id", v1.DelUser)
			//添加用户
			user.POST("/add", v1.AddUser)
			//查看个人资料
			user.GET("/profile", v1.ShowUserProfile)
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
		//Tag模块的路由接口
		tag := authV1.Group("/tag")
		{
			//添加标签
			tag.POST("/add", v1.AddTag)
			//编辑标签
			tag.PUT("/:id", v1.EditTag)
			//删除标签
			tag.DELETE("/:id", v1.DelTag)
		}
		//Comment模块的路由接口
		comment := authV1.Group("comment")
		{
			//编辑评论
			comment.PUT("/:id", v1.EditComment)
			//删除评论
			comment.DELETE("/:id", v1.DelComment)
		}
		//Friends模块的路由接口
		friend := authV1.Group("friend")
		{
			//新增友链
			friend.POST("/add", v1.AddFriends)
			//编辑友链
			friend.PUT("/:id", v1.EditFriends)
			//删除友链
			friend.DELETE("/:id", v1.DelFriends)
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
		pubilcV1.GET("/articles", v1.ShowArticles)
		//查看单个文章内容
		pubilcV1.GET("/article/:id", v1.ShowSingleArticle)

		//Category模块的路由接口
		//查看分类列表
		pubilcV1.GET("/categories", v1.ShowCategories)
		//查看单个分类信息
		pubilcV1.GET("/category/:id", v1.ShowCategoryInfo)
		//查看单个分类下文章
		pubilcV1.GET("/category/articles/:id", v1.ShowCategoryArticles)

		//Tag模块的路由接口
		//查看标签列表
		pubilcV1.GET("/tags", v1.ShowTag)
		//查看单个标签信息
		pubilcV1.GET("/tag/:id", v1.ShowTagInfo)

		//Comment模块的公共路由接口
		//查看评论列表
		pubilcV1.GET("/comments", v1.ShowComments)
		//新增评论
		pubilcV1.POST("/comment/add", v1.AddComment)

		//Friend模块的公共路由接口
		pubilcV1.GET("/friends", v1.ShowFriends)

	}
	err := r.Run(utils.HttpPort)
	if err != nil {
		fmt.Println("路由初始化失败，请检查:", err)
	}
}
