package response

// Response 统一的返回体结构
type Response struct {
	Code    int    `json:"code"`    // 状态码
	Message string `json:"message"` // 消息
	Data    any    `json:"data"`    // 数据
}

// Success 成功返回
func Success(data any) *Response {
	return &Response{
		Code:    0,
		Message: "sucess",
		Data:    data,
	}
}

// Error 错误返回
func Error(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}
