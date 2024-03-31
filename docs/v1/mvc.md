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

