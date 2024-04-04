package model

import (
	"fmt"
	"time"
)

type User struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL" form:"id" json:"id"`
	Name        string    `gorm:"column:name;default:NULL" form:"name" json:"name"`
	Password    string    `gorm:"column:password;default:NULL" form:"password" json:"password"`
	CreatedTime time.Time `gorm:"column:created_time;default:NULL" form:"createdTime" json:"createdTime"`
	UpdatedTime time.Time `gorm:"column:updated_time;default:NULL" form:"updatedTime" json:"updatedTime"`
}

// GetUser 查询用户数据
// https://gorm.io/docs/query.html
func GetUser(user *User) (*User, error) {

	u2 := &User{}

	// no primary key defined, results will be ordered by first field (i.e., `users.name`)
	// SELECT * FROM `users` WHERE name = 'admin' AND password = 'admin123' ORDER BY `users`.`name` LIMIT 1
	tx := Conn.Table("user").Where("name = ? AND password = ?", user.Name, user.Password).First(u2)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u2, nil
}
