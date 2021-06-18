package support

import (
	"github.com/hb0730/auto-sign/application"
	"github.com/hb0730/auto-sign/config"
	"github.com/mritd/logger"
)

var v2ex = application.V2ex{}

type V2ex struct {
	Support
}

func init() {
	logger.Info("[support v2ex] 开始注册 ....")
	v := V2ex{}
	v.Name = "v2ex"
	v.ISupport = v
	Register("v2ex", v)
}

func (v V2ex) DoRun() error {
	logger.Info("[support v2ex] 开始签到 ...")
	yaml := config.ReadYaml()
	cookies := yaml.StringMap(GetV2exYamlKey())
	v2ex.Cookies = cookies
	return v2ex.Start()
}

func GetV2exYamlKey() string {
	return "v2ex.cookies"
}
