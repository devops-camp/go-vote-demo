package logic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutHandler(c *gin.Context) {
	// #1. 清除 session
	_ = FlushSession(c)
	// 重定向到登录页面
	c.Redirect(http.StatusSeeOther, "/")
}
