package logic

import (
	"fmt"
	"net/http"

	"github.com/devops-camp/go-vote-demo/app/model"
	"github.com/devops-camp/go-vote-demo/app/tools"
	"github.com/gin-gonic/gin"
)

func GetLoginHandler(c *gin.Context) {
	// #2. 使用模版文件
	// 完整页面， 没有数据传入
	c.HTML(http.StatusOK, "login.tmpl", nil)
}

func PostLoginHandler(c *gin.Context) {
	// #5 获取表单数据
	user := &model.User{}

	err := c.ShouldBind(user)
	if err != nil {
		c.String(http.StatusBadRequest, "bad request")
		// 显示错误信息之后， 一定要 return 结束后序逻辑
		return
	}

	// # 查询用户
	ret, err := model.GetUser(user)
	if err != nil {

		c.JSON(http.StatusBadRequest, tools.Ecode{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("%v", err),
		})

		return
	}

	// domain 设置为空， 任意域名都可以生效
	// domain 设置不为空， 则仅对特定域名生效
	// c.SetCookie("name", ret.Name, 3600, "/", "", true, false)
	err = SetSession(c, ret.Name, ret.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, tools.Ecode{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("%v", err),
		})
		return
	}

	// 成功后显示用户信息
	// 303: redirect with GET method
	// 307: redirect with Current method (post -> post)
	c.Redirect(http.StatusSeeOther, "/index")
	// c.JSON(http.StatusOK, ret)

}
