package support

import (
	"github.com/hb0730/auto-sign/application"
	"github.com/hb0730/auto-sign/config"
	"github.com/mritd/logger"
)

// Geekhub 支持Geekhub
type Geekhub struct {
	Support
}

var hub = application.GeekHub{}

// init 初始化 注册
func init() {
	logger.Info("[message geekhub] 注册 ....")
	hub := Geekhub{}
	hub.Name = "geekhub"
	hub.ISupport = hub
	Register("geekhub", hub)
}

// DoRun 开始签到
func (g Geekhub) DoRun() error {
	logger.Info("[message geekhub] 开始签到 ....")
	cookies := GetGeekhubYaml()
	hub.Cookies = cookies
	return hub.Start()
}

// GetGeekhubYaml 获取Geekhub yaml配置
func GetGeekhubYaml() map[string]string {
	yaml := config.ReadYaml()
	return yaml.GetStringMapString(GeekhubYamlKey())
}

func GeekhubYamlKey() string {
	return "geekhub.cookies"
}
