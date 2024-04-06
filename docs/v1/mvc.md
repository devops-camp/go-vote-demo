---
title: 5. 拆分为 MVC 目录结构
---

## 1. MVC 目录

```bash
$ tree app
app
  ├── model # Model 数据库相关操作
  ├── view  # View 视图/模版相关操作
  ├── logic # Controller 控制器/代码逻辑相关操作
  ├── router # 路由相关的操作
  └── tools # 公共组件

5 directories, 0 files
```

## 2. 拆分代码到 MVC 目录

```bash
$ tree app
app
  ├── app.go
  ├── logic
  │   ├── login.go
  │   └── ping.go
  ├── model
  │   ├── db.go
  │   └── user.go
  ├── router
  │   └── router.go
  ├── tools
  └── view
    └── login.tmpl

5 directories, 7 files
```

## 3. GORM Model 中的查询的 OrderBy 规则

> https://gorm.io/docs/query.html
>
> The `First` and `Last` methods will find the first and last record (respectively) as ordered by primary key. They only work when a pointer to the destination struct is passed to the methods as argument or when the model is specified using `db.Model()`. Additionally, if no primary key is defined for relevant model, then the model will be ordered by the first field. For example:

```go
// GetUser 查询用户数据
// https://gorm.io/docs/query.html
func GetUser(user *User) (*User, error) {

	u2 := &User{}

	// no primary key defined, results will be ordered by first field (i.e., `user.name`)
	// SELECT * FROM `user` WHERE name = 'admin' AND password = 'admin123' ORDER BY `user`.`name` LIMIT 1
	tx := Conn.Table("user").Where("name = ? AND password = ?", user.Name, user.Password).First(u2)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u2, nil
}
```