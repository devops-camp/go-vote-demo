package model

import "fmt"

type User struct {
	Name     string `form:"name" binding:"required" json:"name"`
	Password string `form:"password" binding:"required" json:"password"`
}

// GetUser 查询用户数据
// https://gorm.io/docs/query.html
func GetUser(user *User) (*User, error) {

	u2 := &User{}

	// no primary key defined, results will be ordered by first field (i.e., `users.name`)
	// SELECT * FROM `users` WHERE name = 'admin' AND password = 'admin123' ORDER BY `users`.`name` LIMIT 1
	tx := Conn.Table("users").Where("name = ? AND password = ?", user.Name, user.Password).First(u2)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u2, nil
}
