package app

import (
	"github.com/devops-camp/go-vote-demo/app/model"
	"github.com/devops-camp/go-vote-demo/app/router"
)

func Start() {
	// 启动数据库连接
	model.NewMysql()
	// 关闭数据库连接
	defer model.Close()

	// 启动 Gin Server
	router.New()
}
