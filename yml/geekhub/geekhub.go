package geekhub

import (
	"auto-sign/geekhub"
	"auto-sign/yml"
)

type Geekhub struct {
	config.AbstractSupport
	Cookies map[string]string `yaml:"cookies,omitempty"`
}

func (Geekhub) Supports(config config.YamlConfig) config.Support {
	// 这里的设置主要解决 *AbstractSupport.Run时nil问题
	// 故儿需要将其重新设置
	g := config.Geekhub
	g.Sub = g
	g.SubName = "geekhub"
	return g
}

//func (g Geekhub) Run() {
//	g.Do(g.Cookies)
//}
func (g Geekhub) DoVoid() {
	g.Do(g.Cookies)
}
func (g Geekhub) Do(config interface{}) {
	if cookies, ok := config.(map[string]string); ok {
		hub := geekhub.Geekhub{Cookies: cookies}
		hub.Do()
	}
}
