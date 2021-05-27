package utils

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

	return "Module :【" + e.Module + "】,Method:【" + e.Method + "】" + "Message:【" + e.Message + "】" + " :" + e.E.Error()
}
