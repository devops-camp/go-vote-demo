package logic

import (
	"net/http"

	"github.com/devops-camp/go-vote-demo/app/tools"
	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	// 这里的 status 不建议直接使用 200，建议使用 http.StatusOK
	// 语义化之后更加直观，不容易出错
	c.JSON(http.StatusOK, tools.Ecode{
		Code:    http.StatusOK,
		Message: "pong",
	})
}
