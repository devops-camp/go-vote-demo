package tools

type Ecode struct {
	Code    int    `json:"code"`    // code, 可以自定错误代码
	Message string `json:"message"` // 消息
	Data    any    `json:"data"`    // 数据
}
