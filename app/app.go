package app

import (
	"github.com/devops-camp/go-vote-demo/app/model"
	"github.com/gin-gonic/gin"
)

func Start() {
	// 启动数据库连接
	model.NewMysql()
	// 关闭数据库连接
	defer model.Close()

	r := gin.Default()

	// #2. 使用 LoadHTMLGlob 加载模板文件
	r.LoadHTMLGlob("app/view/*")

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
