package support

import (
	"fmt"
)

// ReadYaml 读取yaml
type ReadYaml interface {
	//Read 读取yaml
	//Support 返回读取信息
	//error 读取错误
	Read() (ISuperJob, error)
}

//Support 支持
type Support interface {
	DoSupport()
}

//ReadError 读取错误
type ReadError struct {
	Errors []string
}

//Error 错误信息
func (err *ReadError) Error() string {
	return fmt.Sprintf("%v \n", err.Errors)
}
