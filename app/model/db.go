package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// initial 连接数据库
// https://gorm.io/docs/connecting_to_the_database.html
func initial() *gorm.DB {
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
var Conn *gorm.DB

// NewMysql 初始化数据库连接器。 （单例模式）
func NewMysql() {
	if Conn == nil {
		Conn = initial()
	}
}

func Close() {
	db, _ := Conn.DB()
	db.Close()
}
