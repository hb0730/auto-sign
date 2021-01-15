package ld246

import (
	config2 "auto-sign/config"
	"auto-sign/ld246"
	"auto-sign/support"
	"auto-sign/util"
)

type Ld246 struct {
	support.AbstractSupport
	User User `json:"user"`
}
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l Ld246) Read() (support.ISuperJob, error) {
	provider, err := config2.ReadYaml()
	if err != nil {
		util.ErrorF("read support file error, %v \n", err)
		return l, &support.ReadError{Errors: []string{"读取yaml文件失败失败:message", err.Error()}}
	}
	var result Ld246
	_ = provider.Get("ld246").Populate(&result)
	result.Sub = result
	result.SubName = "ld246"
	return result, nil
}
func (ld Ld246) DoSupport() {
	l := ld246.LD{Username: ld.User.Username, Password: ld.User.Password}
	l.Do()

}
