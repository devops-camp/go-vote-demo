---
title: 5. 拆分为 MVC 目录结构
---

## 1. 拆分 MVC 目录

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
