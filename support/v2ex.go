package support

import (
	"github.com/hb0730/auto-sign/application"
	"github.com/hb0730/auto-sign/config"
	"github.com/hb0730/auto-sign/utils"
)

var v2ex = application.V2ex{}

type V2ex struct {
	Support
}

func init() {
	utils.Info("v2ex 开始注册")
	v := V2ex{}
	v.ISupport = v
	Register("v2ex", v)
}

func (v V2ex) DoRun() error {
	utils.Info("v2ex 开始签到")
	yaml := config.ReadYaml()
	cookies := yaml.GetStringMapString("v2ex.cookies")
	v2ex.Cookies = cookies
	return v2ex.Start()
}
