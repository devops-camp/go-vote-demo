---
title: 08. 设计所有表
date: "2024-04-04T08:11:23Z"
---

1. 用户表
    + 管理用户列表
2. 投票表
    + 管理投票列表
    + 管理投票与其创建者的关系
3. 投票选项表
    + 管理选项
    + 管理选项与其绑定投票的关系
4. 选项-用户表
    + 管理用户与投票选项的关系。
    + 避免用户重复投票。

![](./design-tables.png)

```sql
-- 用户表
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL,
  `password` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL,
  `created_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updated_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 投票表
-- 绑定投票与其创建用户的关系
CREATE TABLE `vote` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `type` int DEFAULT NULL COMMENT '0 for single, 1 for multiple choice',
  `status` int DEFAULT NULL COMMENT '0 for normal, 1 for expired',
  `user_id` bigint DEFAULT NULL COMMENT 'who created',
  `expired_in` bigint DEFAULT NULL,
  `created_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updated_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 投票选项表
-- 绑定投票选项表与投票表关系
CREATE TABLE `vote_opt` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '选项名称',
  `count` int DEFAULT NULL COMMENT '选项投票计数器',
  `vote_id` bigint DEFAULT NULL COMMENT '绑定到的投票单',
  `created_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updated_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 投票选项用户关系表
-- 防止用户重复投票
CREATE TABLE `vote_opt_user` (
  `id` bigint NOT NULL,
  `vote_id` bigint DEFAULT NULL COMMENT '表单ID',
  `user_id` bigint DEFAULT NULL COMMENT '投票用户ID',
  `vote_opt_id` bigint DEFAULT NULL COMMENT '选项ID',
  `created_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updated_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
```


## 2. 将 SQL 转换为 ORM 模型

使用在线工具， 将 SQL 转换为 ORM 模型。

> http://www.gotool.top/handlesql/sql2gorm

这我个人觉得还缺少一个 **软删除** 字段 `DeletedTime`。 用于标记字段是否删除， 以及删除时间。  

```go
package model

import "time"

type Vote struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Title       string    `gorm:"column:title;default:NULL"`
	Type        int32     `gorm:"column:type;default:NULL;comment:'0 for single, 1 for multiple choice'"`
	Status      int32     `gorm:"column:status;default:NULL;comment:'0 for normal, 1 for expired'"`
	UserId      int64     `gorm:"column:user_id;default:NULL;comment:'who created'"`
	ExpiredIn   int64     `gorm:"column:expired_in;default:NULL"`
	CreatedTime time.Time `gorm:"column:created_time;default:NULL"`
	UpdatedTime time.Time `gorm:"column:updated_time;default:NULL"`
}

func (v *Vote) TableName() string {
	return "vote"
}

type VoteOpt struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Name        string    `gorm:"column:name;default:NULL;comment:'选项名称'"`
	Count       int32     `gorm:"column:count;default:NULL;comment:'选项投票计数器'"`
	VoteId      int64     `gorm:"column:vote_id;default:NULL;comment:'绑定到的投票单'"`
	CreatedTime time.Time `gorm:"column:created_time;default:NULL"`
	UpdatedTime time.Time `gorm:"column:updated_time;default:NULL"`
}

func (v *VoteOpt) TableName() string {
	return "vote_opt"
}

type VoteOptUser struct {
	Id          int64     `gorm:"column:id;primary_key;NOT NULL"`
	VoteId      int64     `gorm:"column:vote_id;default:NULL;comment:'表单ID'"`
	UserId      int64     `gorm:"column:user_id;default:NULL;comment:'投票用户ID'"`
	VoteOptId   int64     `gorm:"column:vote_opt_id;default:NULL;comment:'选项ID'"`
	CreatedTime time.Time `gorm:"column:created_time;default:NULL"`
	UpdatedTime time.Time `gorm:"column:updated_time;default:NULL"`
}

func (v *VoteOptUser) TableName() string {
	return "vote_opt_user"
}
```
