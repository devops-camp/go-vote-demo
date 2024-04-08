package logic

import (
	"fmt"
	"net/http"

	"github.com/devops-camp/go-vote-demo/app/model"
	"github.com/devops-camp/go-vote-demo/app/tools"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func IndexLogin(c *gin.Context) {

	name, _ := c.Cookie("name")

	votes, err := model.GetVotes()
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.Ecode{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})

		return
	}

	data := map[string]any{
		"Name":  name,
		"Votes": votes,
	}
	c.HTML(200, "index-login.tmpl", data)
}

// IndexLoginCheckerMiddleware 检查 cookie 则跳转到登录页面
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

func IndexLoginCheckSessionMiddleware(c *gin.Context) {
	var name string
	var id int64

	values := GetSession(c)
	if v, ok := values["name"]; ok {
		name = v.(string)
	}
	if v, ok := values["id"]; ok {
		id = v.(int64)
	}

	if name == "" || id <= 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, tools.Ecode{
			Code:    http.StatusUnauthorized,
			Message: "Session Failed: Unauthorized",
		})
		return
	}

	fmt.Println("name: ", name, "id", id)
}
