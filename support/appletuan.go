package support

import (
	"github.com/hb0730/auto-sign/application"
	"github.com/hb0730/auto-sign/config"
	"github.com/mritd/logger"
)

type AppleTuan struct {
	Support
}

var apple = application.AppleTuan{}

func init() {
	logger.Info("[support appletuan] 开始注册 ....")
	tuan := AppleTuan{}
	tuan.ISupport = tuan
	tuan.Name = "苹果团"
	Register("appletuan", tuan)
}

func (tuan AppleTuan) DoRun() error {
	logger.Info("[support appletuan] 开始签到 ")
	cookies := GetAppleTuanYaml()
	apple.Cookies = cookies
	return apple.Start()
}

func GetAppleTuanYamlKey() string {
	return "appletuan.cookies"
}

func GetAppleTuanYaml() map[string]string {
	yaml := config.ReadYaml()
	return yaml.StringMap(GetAppleTuanYamlKey())
}
