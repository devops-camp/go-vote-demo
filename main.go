package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	// 健康检查接口
	r.GET("/ping", pingHandler)

	// #2. 使用 LoadHTMLGlob 加载模板文件
	r.LoadHTMLGlob("tmpl/*")

	// Login GET
	r.GET("/login", getLoginHandler)
	// #5. 获取表单数据
	r.POST("/login", postLoginHandler)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func pingHandler(c *gin.Context) {
	// 这里的 status 不建议直接使用 200，建议使用 http.StatusOK
	// 语义化之后更加直观，不容易出错
	c.String(http.StatusOK, "pong")
}

func getLoginHandler(c *gin.Context) {
	// #2. 使用模版文件
	// 完整页面， 没有数据传入
	c.HTML(http.StatusOK, "login.tmpl", nil)
}

func postLoginHandler(c *gin.Context) {
	// #5 获取表单数据
	user := &User{}

	err := c.ShouldBind(user)
	if err != nil {
		c.String(http.StatusBadRequest, "bad request")
		// 显示错误信息之后， 一定要 return 结束后序逻辑
		return
	}

	// #7. 连接数据库, 查询用户
	// https://gorm.io/docs/query.html
	tx := GlobalConn.Table("users").Where("name = ? AND password = ?", user.Name, user.Password).First(user)
	if tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "user not found",
			"error": fmt.Sprintf("%v", tx.Error),
		})

		return
	}

	// 成功后显示用户信息
	c.JSON(http.StatusOK, user)
}

type User struct {
	Name     string `form:"name" binding:"required" json:"name"`
	Password string `form:"password" binding:"required" json:"password"`
}

// DBConn 连接数据库
// #7. 连接数据库
// https://gorm.io/docs/connecting_to_the_database.html
func DBConn() *gorm.DB {
	dsnfmt := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := fmt.Sprintf(dsnfmt,
		"root", "Mysql12345",
		"127.0.0.1", 3306,
		"vote")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db

}

// 定义全局 DB 连接
var GlobalConn *gorm.DB

// 初始化函数， 在 main 函数之前执行
func init() {
	if GlobalConn == nil {
		GlobalConn = DBConn()
	}
}
