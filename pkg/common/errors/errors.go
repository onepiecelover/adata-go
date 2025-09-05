// Package errors 定义项目错误类型
package errors

import "fmt"

// ADataError adata自定义错误类型
type ADataError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}

func (e *ADataError) Error() string {
	if e.Detail != "" {
		return fmt.Sprintf("[%d] %s: %s", e.Code, e.Message, e.Detail)
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// 预定义错误
var (
	// ErrInvalidStockCode 无效股票代码
	ErrInvalidStockCode = &ADataError{
		Code:    10001,
		Message: "无效的股票代码",
	}

	// ErrInvalidDateFormat 无效日期格式
	ErrInvalidDateFormat = &ADataError{
		Code:    10002,
		Message: "无效的日期格式",
	}

	// ErrRequestFailed 请求失败
	ErrRequestFailed = &ADataError{
		Code:    20001,
		Message: "请求失败",
	}

	// ErrParseResponseFailed 解析响应失败
	ErrParseResponseFailed = &ADataError{
		Code:    20002,
		Message: "解析响应失败",
	}

	// ErrNoDataFound 未找到数据
	ErrNoDataFound = &ADataError{
		Code:    30001,
		Message: "未找到数据",
	}

	// ErrDataSourceUnavailable 数据源不可用
	ErrDataSourceUnavailable = &ADataError{
		Code:    30002,
		Message: "数据源不可用",
	}
)

// NewADataError 创建自定义错误
func NewADataError(code int, message, detail string) *ADataError {
	return &ADataError{
		Code:    code,
		Message: message,
		Detail:  detail,
	}
}

// WrapError 包装错误
func WrapError(err error, message string) *ADataError {
	return &ADataError{
		Code:    99999,
		Message: message,
		Detail:  err.Error(),
	}
}
