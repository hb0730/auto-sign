package support

import (
	"github.com/hb0730/auto-sign/application"
	"github.com/hb0730/auto-sign/config"
	"github.com/hb0730/auto-sign/utils"
)

type AppleTuan struct {
	Support
}

var apple = application.AppleTuan{}

func init() {
	utils.Info("appleTuan 开始注册")
	tuan := AppleTuan{}
	tuan.ISupport = tuan
	Register("appletuan", tuan)
}

func (tuan AppleTuan) DoRun() error {
	utils.Info("appletuan 开始签到")
	yaml := config.ReadYaml()
	cookies := yaml.GetStringMapString("appletuan.cookies")

	apple.Cookies = cookies
	return apple.Start()
}
