package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"goblog/config"
	"goblog/controller"
	"goblog/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 初始化 session 中间件，使用 cookie 存储
	store := cookie.NewStore([]byte(config.GetSessionkey()))
	r.Use(sessions.Sessions("mysession", store))
	r.Use(middleware.CheckLoginMiddleware)

	// 设置页面模板目录
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	// 设置路由
	r.GET("/", controller.ShowPosts)            // 首页，显示所有帖子
	r.GET("/posts/:id", controller.ShowPost)    // 查看单个帖子
	r.GET("/posts/new", controller.NewPostForm) // 显示新建帖子表单
	r.POST("/posts", controller.CreatePost)     // 创建新帖子

	// 登录注册
	r.GET("/register", controller.ShowRegisterForm)
	r.POST("/register", controller.Register)
	r.GET("/login", controller.ShowLoginForm)
	r.POST("/login", controller.Login)

	// 搜索文章
	r.GET("/search", controller.SearchPosts)
	// 退出登陆
	r.GET("/logout", controller.Logout)
	// 个人主页
	r.GET("/user/:userName", controller.UserHome)
	// 排行榜 (贴子数)
	r.GET("/rank", controller.RankByPosts)
	return r
}
