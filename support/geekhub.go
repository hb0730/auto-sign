package support

import (
	"github.com/hb0730/auto-sign/application"
	"github.com/hb0730/auto-sign/config"
	"github.com/hb0730/auto-sign/utils"
)

// Geekhub 支持Geekhub
type Geekhub struct {
	Support
}

var hub = application.GeekHub{}

// init 初始化 注册
func init() {
	utils.Info("geekhub 注册 ....")
	hub := Geekhub{}
	hub.Name = "geekhub"
	hub.ISupport = hub
	Register("geekhub", hub)
}

// DoRun 开始签到
func (g Geekhub) DoRun() error {
	utils.Info("geekhub 开始签到 ....")
	yaml := config.ReadYaml()
	cookies := yaml.GetStringMapString("geekhub.cookies")
	hub.Cookies = cookies
	return hub.Start()
}
