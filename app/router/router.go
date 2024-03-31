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

	// Login GET
	r.GET("/login", logic.GetLoginHandler)
	// 获取表单数据
	r.POST("/login", logic.PostLoginHandler)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}