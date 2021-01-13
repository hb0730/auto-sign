package appletuan

import (
	"auto-sign/appletuan"
	config2 "auto-sign/config"
	"auto-sign/support"
	"auto-sign/util"
)

type AppleTuan struct {
	support.AbstractSupport
	Cookies map[string]string `json:"cookies"`
}

func (t AppleTuan) Read() (support.ISuperJob, error) {
	provider, err := config2.ReadYaml()
	if err != nil {
		util.ErrorF("read support file error, %v \n", err)
		return t, &support.ReadError{Errors: []string{"读取yaml文件失败失败:message", err.Error()}}
	}
	var result AppleTuan
	_ = provider.Get("appletuan").Populate(&result)
	result.Sub = result
	result.SubName = "appletuan"
	return result, nil
}

func (t AppleTuan) DoSupport() {
	util.Info("appletuan doSupport start ....")
	tuan := appletuan.AppleTuan{Cookies: t.Cookies}
	tuan.Do()
}
