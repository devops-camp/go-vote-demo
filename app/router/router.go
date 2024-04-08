package router

import (
	"github.com/devops-camp/go-vote-demo/app/logic"
	"github.com/gin-gonic/gin"
)

func New() {
	r := gin.Default()

	// 健康检查接口
	r.GET("/ping", logic.PingHandler)

	// 使用 LoadHTMLGlob 加载模板文件
	// 虽然在 app/router/router.go 中创建的 gin Server
	// 但是在引用的模板文件中，需要使用相对于 main.go 的相对路径
	r.LoadHTMLGlob("app/view/*")

	// 首页
	r.GET("/", logic.Index)

	{
		authorized := r.Group("")
		// 使用中间件， 检查cookie
		// authorized.Use(logic.IndexLoginCheckerMiddleware)
		authorized.Use(logic.IndexLoginCheckSessionMiddleware)

		// 登录后首页
		authorized.GET("/index", logic.IndexLogin)
		authorized.GET("/vote", logic.GetVoteHandler)
		authorized.POST("/vote", logic.PostVoteHandler)
	}

	// Login GET
	r.GET("/login", logic.GetLoginHandler)
	// 获取表单数据
	r.POST("/login", logic.PostLoginHandler)
	r.GET("/logout", logic.LogoutHandler)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
