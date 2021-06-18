package support

import (
	"github.com/hb0730/auto-sign/application"
	"github.com/hb0730/auto-sign/config"
	"github.com/mritd/logger"
)

type ChinaG struct {
	Support
}

var gg = application.ChinaG{}

func init() {
	logger.Info("[support chinaG] 开始注册 ... ")
	g := ChinaG{}
	g.ISupport = g
	g.Name = "chinaG"
	Register("chinag", g)
}

func (g ChinaG) DoRun() error {
	logger.Info("[support chinaG] 开始签到 ...")
	yaml := config.ReadYaml()
	u := yaml.StringMap(GetChinaGYamlKey())
	var username = u["username"]
	var password = u["password"]
	gg.Username = username
	gg.Password = password
	return gg.Start()
}

func GetChinaGYamlKey() string {
	return "chinaG.user"
}
