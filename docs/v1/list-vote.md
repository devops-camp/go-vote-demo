---
title: 09. 列出所有投票并撰写单元测试
date: "2024-12-06T16:12Z"
---

## 1. 列出所有投票结果

在 `/model/vote.go` 中添加以下代码：

```go
// GetVotes 获取投票列表
func GetVotes() ([]Vote, error) {
	votes := make([]Vote, 0)
	tx := Conn.Table("vote").Find(&votes)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return votes, nil
}
```

## 2. 插入数据

在 vote 表中插入一些数据：

```sql
INSERT INTO `vote` (`id`, `title`, `type`, `status`, `user_id`, `expired_in`, `created_time`, `updated_time`) VALUES
(1, '今天晚上吃什么', 0, 0, 1, 86400, '2024-04-06 08:03:01', '2024-04-06 08:03:01'),
(2, '你会用vscode吗', 0, 0, 2, 86400, '2024-04-06 08:03:01', '2024-04-06 08:03:01');
```

## 3. 撰写简单的单元测试

创建 `/model/vote_test.go`, 编写一个简单的单元测试：

```go
func init() {
	NewMysql()
}

// TestGetVotes 测试获取投票列表
func TestGetVotes(t *testing.T) {
	votes, err := GetVotes()

	if err != nil {
		t.Errorf("GetVotes() failed: %v", err)
	}

	//t.Logf("votes: %v", votes)
	for _, v := range votes {
		t.Logf("vote: %v", v)
	}
}
```

测试结果如下

```log
=== RUN   TestGetVotes
    vote_test.go:20: vote: {1 今天晚上吃什么 0 0 1 86400 2024-04-06 08:03:01 +0800 CST 2024-04-06 08:03:01 +0800 CST}
    vote_test.go:20: vote: {2 你会用vscode吗 0 0 2 86400 2024-04-06 08:03:01 +0800 CST 2024-04-06 08:03:01 +0800 CST}
```
