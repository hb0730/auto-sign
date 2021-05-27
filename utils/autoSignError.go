package utils

import "strings"

// AutoSignError auto-sign异常
type AutoSignError struct {
	//Module 模块
	Module string
	// Method 方法
	Method string
	// Message 信息
	Message string
	// E 异常
	E error
}

// AsError 提供外部使用
var AsError error = &AutoSignError{}

func (e AutoSignError) Error() string {
	b := strings.Builder{}
	b.WriteString("Module:【 ")
	b.WriteString(e.Module)
	b.WriteString("】, Method:【 ")
	b.WriteString(e.Method)
	b.WriteString(" 】，Message:【 ")
	b.WriteString(e.Message)
	b.WriteString(" 】")
	if e.E != nil {
		b.WriteString(",Error:【 ")
		b.WriteString(e.E.Error())
		b.WriteString(" 】")
	}
	return b.String()
}
