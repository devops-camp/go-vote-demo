---
title: 10. 渲染数据到模板
date: "2024-38-06T17:38Z"
---

## 1. 渲染数据到登录后首页

> https://www.bilibili.com/video/BV1AZ42117JH/

1. 在 `logic/index.go` 页面中从 model.vote 中数据所有投票数据。

```go
func IndexLogin(c *gin.Context) {

	name, _ := c.Cookie("name")

	// 获取所有投票数据
	votes, err := model.GetVotes()
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.Ecode{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// 添加到 data 中进行渲染
	data := map[string]any{
		"Name":  name,
		"Votes": votes,
	}
	c.HTML(200, "index-login.tmpl", data)
}
```

2. 在 `view/index-login.tmpl` 中添加投票数据渲染模版

```html
<h3>投票列表</h3>
{{ range $key, $value := .Votes }}
<ul>
    <li>{{ $key}} - {{ $value.Title }}</li>
</ul>
{{ end }}
```

![](./get-votes.png)


## 2. 根据ID获取投票信息并渲染

> https://www.bilibili.com/video/BV1at421t7e7/

1. 在 `model/vote.go` 中添加根据ID获取投票信息的方法

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

2. 在 `view/vote.html` 中创建投票详情模版

```html
<!-- vote 信息-->
<div class="vote">
    <h2>{{ .Title }}</h2>
    <span>状态: {{ .Status }}</span>
    <span>创建用户: {{ .UserId }}</span>
    <span>过期时间: {{ .ExpiredIn }}</span>
    <span>创建时间: {{ .CreatedTime }}</span>
</div>
```

3. 在 `logic/vote.go`  添加根据查询Vote数据并渲染页面

```go
func GetVoteHandler(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, tools.EcodeBadRequest("id is required"))

		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.EcodeBadRequest("id is invalid"))
		return
	}

	vote, err := model.GetVote(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, tools.EcodeBadRequest(err.Error()))
		return
	}

	//c.JSON(http.StatusOK, vote)
	c.HTML(http.StatusOK, "vote.html", vote)
}
```

![](vote-detail.png)

