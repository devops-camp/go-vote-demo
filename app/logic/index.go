package logic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func IndexLogin(c *gin.Context) {

	name, _ := c.Cookie("name")
	data := map[string]string{
		"Name": name,
	}
	c.HTML(200, "index-login.tmpl", data)

}

// IndexLoginCheckerMiddleware 检查 cookie 则跳转到登录也
func IndexLoginCheckerMiddleware(c *gin.Context) {

	name, err := c.Cookie("name")
	if err != nil || name == "" {
		// 跳转到登录页面
		c.Redirect(http.StatusSeeOther, "/login")
		// c.Abort()
		return
	}

	c.Next()
}
