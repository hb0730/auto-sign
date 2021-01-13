package v2ex

import (
	config2 "auto-sign/config"
	"auto-sign/support"
	"auto-sign/util"
	"auto-sign/v2ex"
)

//V2ex http://v2ex.com
type V2ex struct {
	support.AbstractSupport
	Cookies map[string]string `json:"cookies"`
}

func (v V2ex) DoSupport() {
	ex := v2ex.V2ex{Cookies: v.Cookies}
	ex.Do()
}
func (v V2ex) Read() (support.ISuperJob, error) {
	provider, err := config2.ReadYaml()
	if err != nil {
		util.ErrorF("read support file error, %v \n", err)
		return v, &support.ReadError{Errors: []string{"读取yaml文件失败失败:message", err.Error()}}
	}
	var result V2ex
	_ = provider.Get("v2ex").Populate(&result)
	result.Sub = result
	result.SubName = "ld246"
	return result, nil
}
