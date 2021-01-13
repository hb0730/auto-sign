package geekhub

import (
	config2 "auto-sign/config"
	"auto-sign/geekhub"
	"auto-sign/support"
	"auto-sign/util"
)

//Geekhub https://geekhub.com
type Geekhub struct {
	support.AbstractSupport
	Cookies map[string]string `json:"cookies"`
}

func (Geekhub) Read() (support.ISuperJob, error) {
	provider, err := config2.ReadYaml()
	if err != nil {
		util.ErrorF("read support file error, %v \n", err)
		return Geekhub{}, &support.ReadError{Errors: []string{"格式转换错误:message", err.Error()}}
	}
	var result Geekhub
	err = provider.Get("geekhub").Populate(&result)
	if err != nil {
		util.ErrorF("read support file error, %v \n", err)
		return Geekhub{}, &support.ReadError{Errors: []string{"获取geekhub配置失败:message", err.Error()}}
	}
	result.Sub = result
	result.SubName = "Geekhub"
	return result, err
}

func (g Geekhub) DoSupport() {
	util.Info("geekhub doSupport start ....")
	hub := geekhub.Geekhub{Cookies: g.Cookies}
	hub.Do()
}
