package cron

import (
	config2 "auto-sign/config"
	"auto-sign/support"
	"auto-sign/util"
)

//Cron 表达式
type Cron struct {
	support.Support
	Cron map[string]string
}

//Read 读取支持的表达式
func Read() (Cron, error) {
	provider, err := config2.ReadYaml()
	if err != nil {
		util.ErrorF("read support file error, %v \n", err)
		return Cron{}, &support.ReadError{Errors: []string{"读取yaml文件失败失败:message", err.Error()}}
	}
	var result map[string]string
	err = provider.Get("cron").Populate(&result)
	if err != nil {
		util.ErrorF("read support file error, %v \n", err)
		return Cron{}, &support.ReadError{Errors: []string{"格式转换错误:message", err.Error()}}
	}
	return Cron{Cron: result}, err
}
