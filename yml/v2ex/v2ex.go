package v2ex

import (
	v2ex2 "auto-sign/v2ex"
	"auto-sign/yml"
)

type V2ex struct {
	config.AbstractSupport
	Cookies map[string]string `yaml:"cookies,omitempty"`
}

func (v2ex V2ex) Do(config interface{}) {
	if cookies, ok := config.(map[string]string); ok {
		v2 := v2ex2.V2ex{Cookies: cookies}
		v2.Do()
	}
}

//func (v2 V2ex) Run() {
//	v2.Do(v2.Cookies)
//}

func (v2 V2ex) DoVoid() {
	v2.Do(v2.Cookies)
}

func (v2 V2ex) Supports(config config.YamlConfig) config.Support {
	// 这里的设置主要解决 *AbstractSupport.Run时nil问题
	// 故儿需要将其重新设置
	c := config.V2ex
	c.Sub = c
	c.SubName = "v2ex"
	return c
}
