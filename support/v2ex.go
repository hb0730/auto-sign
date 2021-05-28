package support

import (
	"github.com/hb0730/auto-sign/application"
	"github.com/hb0730/auto-sign/config"
	"github.com/hb0730/auto-sign/utils"
	"strings"
)

var v2ex = application.V2ex{}

type V2ex struct {
	Support
}

func init() {
	utils.Info("[v2ex] 开始注册 ....")
	v := V2ex{}
	v.Name = "v2ex"
	v.ISupport = v
	Register("v2ex", v)
}

func (v V2ex) DoRun() error {
	utils.Info("[v2ex] 开始签到 ")
	yaml := config.ReadYaml()
	cookies := yaml.GetStringMapString("v2ex.cookies")
	var cookie = make(map[string]string, 0)
	for k, v := range cookies {
		cookie[strings.ToUpper(k)] = v
	}
	v2ex.Cookies = cookie
	return v2ex.Start()
}
