package model

type User struct {
	Name     string `form:"name" binding:"required" json:"name"`
	Password string `form:"password" binding:"required" json:"password"`
}
