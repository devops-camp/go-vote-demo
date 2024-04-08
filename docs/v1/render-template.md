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

![](./vote-detail.png)

## 3. 关联 Index 和 Vote 页面

在 `view/index-login.tmpl` 中添加跳转链接

1. 使用 a 标签跳转到 vote 页面
2. 使用 `{{ $value.Id}}` 传递参数

```html
<h3>投票列表</h3>
{{ range $key, $value := .Votes }}
<ul>
    <li>
        <!-- 使用 a 标签传递跳转页面。  -->
        <a href="/vote?id={{ $value.Id}} ">{{ $key}} - {{ $value.Title }}</a>
    </li>
</ul>
{{ end }}
```

![](./get-votes-link.png)


## 4. 展示 VoteOpts 选项

1. 在 `model/vote.go` 中添加获取选项信息的方法

```go
// GetVoteOptsByVoteId 根据投票 ID 查询选项
func GetVoteOptsByVoteId(voteId int64) ([]VoteOpt, error) {
	voteOpts := make([]VoteOpt, 0)
	tx := Conn.Table("vote_opt").Where("vote_id = ?", voteId).Find(&voteOpts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return voteOpts, nil
}
```

2. 在 `logic/vote.go` 中获取选项信息并渲染到页面。

> 注意: 由于同时返回了 Vote 和 VoteOpts 信息。 因此使用 `map[string]any` 作为数据容器。 对应的模版也需要进行更改。，

```go
func GetVoteHandler(c *gin.Context) {

	vote, err := model.GetVote(id)
	opts, err := model.GetVoteOptsByVoteId(id)

	// 使用 data 组合 vote 和 vote_opt 数据
	data := map[string]any{
		"Vote": vote,
		"Opts": opts,
	}

	c.HTML(http.StatusOK, "vote.html", data)
}
```



2. 在 `view/vote.html` 中， 添加选项信息展示

```html
<!-- vote 信息-->
<!-- 修改数据结构字段， 增加 .Vote -->
<div class="vote">
    <h2>{{ .Vote.Title }}</h2>
    <span>状态: {{ .Vote.Status }}</span>
    <span>创建用户: {{ .Vote.UserId }}</span>
    <span>过期时间: {{ .Vote.ExpiredIn }}</span>
    <span>创建时间: {{ .Vote.CreatedTime }}</span>
</div>

<!-- 新增 VoteOpt 信息-->
<!-- vote_opt 信息-->
<div class="vote-opt">
    <ul>
        {{ range $key, $value := .Opts }}
        <li>{{ $value.Id}} - {{ $value.Name }} - {{ $value.Count }}</li>
        {{ end }}
    </ul>
</div>
```

![](./vote-detail-with-opt.png)


## 5. Post VoteOpts 数据到数据库并重新渲染页面


1. 在 `model/vote.go` 中添加更新方法

```go
// UpdateVoteCount 更新 VoteOpt 表计数器
func UpdateVoteCount(id int64, voteId int64) error {
	tx := Conn.Table("vote_opt").
		Where("id = ? AND vote_id = ?", id, voteId).
		Update("count", gorm.Expr("count + ?", 1))
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
```

同时在数据库中插入预设字段

```sql
INSERT INTO `vote_opt` (`id`, `name`, `count`, `vote_id`, `created_time`, `updated_time`) VALUES
 (1, '红烧肉', 0, 1, '2024-04-06 15:07:30', '2024-04-06 15:07:30'),
 (2, '回锅肉', 0, 1, '2024-04-06 15:07:30', '2024-04-06 15:07:30'),
 (3, '东坡肉', 0, 1, '2024-04-06 15:07:30', '2024-04-06 15:07:30'),
 (4, '会用', 3, 2, '2024-04-07 07:45:18', '2024-04-07 07:45:18'),
 (5, '会用一点', 4, 2, '2024-04-07 07:45:21', '2024-04-07 07:45:21'),
 (6, '不会用', 4, 2, '2024-04-07 07:45:21', '2024-04-07 07:45:21');
```


2. 更新 `view/vote.html` 显示更多选项信息
   1. 使用 `input:checkbox` 作为多选框。 (Todo: 这里是否使用多选框， 应该更具数据库中 Vote 的 Type 字段值确定)
   2. 设置一个隐藏 `input` 用于提交 vote_id 信息。
   3. 使用 `form` 表单提交数据

```html
<!-- vote 选项信息-->
<div class="vote-opt">
    <form action="/vote" method="post">
        <!-- 1. 向服务端提交 vote_id, 2. 但不在页面上显示-->
        <input hidden="hidden" name="vote_id"  type="text" value="{{ .Vote.Id }}" />
        {{ range $key, $value := .Opts }}

        <div>
            <!-- name 统一使用 opts, 服务端使用 数组/切片 接收   -->
            <input type="checkbox" name="opts" id="opt{{$value.Id}}" value="{{$value.Id}}" />
            <label for="opt{{$value.Id}}">{{$value.Name}} --- {{ $value.Count }}次</label>
        </div>
        {{ end }}
        <input type="submit" value="Submit" />
    </form>
</div>
```

请求 Payload 为

```http
POST http://127.0.0.1:8080/vote
Content-Type: application/x-www-form-urlencoded

vote_id=2&opts=5&opts=6
```

3. 在 `logic/vote.go` 中添加处理提交数据的方法

```go
type PostVoteParams struct {
	VoteId int64   `form:"vote_id" json:"vote_id" binding:"required"`
	// 使用 切片 接收
	Opts   []int64 `form:"opts" json:"opts" binding:"required"`
}

func PostVoteHandler(c *gin.Context) {
	p := &PostVoteParams{}
	if err := c.ShouldBind(p); err != nil {
		c.JSON(http.StatusBadRequest, tools.EcodeBadRequest(err.Error()))
		return
	}

	for _, id := range p.Opts {
		err := model.UpdateVoteCount(id, p.VoteId)
		if err != nil {
			panic(err)
		}
	}

	// 重定向到 GET vote 页面
	c.Redirect(http.StatusSeeOther, "/vote?id="+strconv.FormatInt(p.VoteId, 10))
}
```

4. 在 `router/router.go` 中添加路由

```go
authorized.POST("/vote", logic.PostVoteHandler)
```

![](./post-vote-opts.png)
