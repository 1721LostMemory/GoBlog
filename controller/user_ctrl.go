package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"goblog/service"
	"net/http"
)

func ShowRegisterForm(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func Register(c *gin.Context) {
	var input struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
		Gender   string `form:"gender"`
		Age      uint   `form:"age"`
		School   string `form:"school"`
		Email    string `form:"email"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "参数错误"})
		return
	}

	// 调用 Service 层注册用户
	err := service.RegisterUser(input.Username, input.Password, input.Gender, input.School, input.Email, input.Age)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{"error": err.Error()})
		return
	}

	// 注册成功，重定向到登录页面
	c.Redirect(http.StatusFound, "/login")
}

func ShowLoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func Login(c *gin.Context) {
	var input struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "参数错误"})
		return
	}

	user, err := service.FindUserByNameAndPwd(input.Username, input.Password)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "用户不存在"})
		return
	}

	// 登录成功，设置 Session
	session := sessions.Default(c)
	session.Set("user", user.Username) // 存储用户名
	session.Save()                     // 保存 Session 数据

	c.Redirect(http.StatusFound, "/")
}

func Logout(c *gin.Context) {
	fmt.Println("Logout~")
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.HTML(http.StatusOK, "index.html", gin.H{"message": "注销成功"})
}

func UserHome(c *gin.Context) {
	userName := c.Param("userName")

	// 查找用户信息
	user, err := service.FindUserByName(userName)
	if err != nil {
		// 如果找不到用户，返回 404 错误
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// 获取用户的帖子
	posts, err := service.GetPostByAuthor(userName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch posts"})
		return
	}

	// 获取当前用户信息
	curUser, _ := c.Get("user")
	if curUser == nil {
		// 如果没有登录，跳转到登录页面
		c.HTML(http.StatusOK, "login.html", gin.H{
			"Title": "登录",
		})
	} else if curUser != user.Username {
		// 如果当前用户与目标用户不一致，返回 403 错误
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to view this user's profile"})
		return
	} else {
		// 如果当前用户与目标用户一致，渲染个人主页
		c.HTML(http.StatusOK, "home.html", gin.H{
			"userName": user.Username,
			"gender":   user.Gender,
			"age":      user.Age,
			"school":   user.School,
			"email":    user.Email,
			"posts":    posts,
		})
	}
}
