package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CheckLoginMiddleware(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")

	var isLogin bool
	if user == nil {
		isLogin = false
	} else {
		isLogin = true
		c.Set("user", user)
	}
	//fmt.Println("Is user logged in:", isLogin)
	//fmt.Println("User:", user)

	c.Set("isLogin", isLogin)

	c.Next()
}
