package tools

import "net/http"

type Ecode struct {
	Code    int    `json:"code"`    // code, 可以自定错误代码
	Message string `json:"message"` // 消息
	Data    any    `json:"data"`    // 数据
}

func EcodeBadRequest(msg string) Ecode {
	return Ecode{
		Code:    http.StatusBadRequest,
		Message: msg,
	}
}

func EcodeSuccess(data any) Ecode {
	return Ecode{
		Code: http.StatusOK,
		Data: data,
	}
}
