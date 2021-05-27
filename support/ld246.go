package support

import (
	"github.com/hb0730/auto-sign/application"
	"github.com/hb0730/auto-sign/config"
	"github.com/hb0730/auto-sign/utils"
)

var ld246 = application.Ld246{}

type Ld246 struct {
	Support
}

func init() {
	utils.Info("ld246开始注册....")
	ld := Ld246{}
	ld.ISupport = ld
	Register("ld246", ld)
}

func (ld Ld246) DoRun() error {
	utils.Info("ld246 开始签到...")
	yaml := config.ReadYaml()
	user := yaml.GetStringMapString("ld246.user")
	ld246.Username = user["username"]
	ld246.Password = user["password"]
	return ld246.Start()
}
