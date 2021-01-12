package appletuan

import (
	autoAppletuan "auto-sign/appletuan"
	"auto-sign/yml"
)

// AppleTuan  https://appletuan.com/
type AppleTuan struct {
	config.AbstractSupport
	//Cookies 用于签到
	Cookies map[string]string `yaml:"cookies,omitempty"`
}

//Supports 所支持的，返回具体类型
func (AppleTuan) Supports(config config.YamlConfig) config.Support {
	// 这里的设置主要解决 *AbstractSupport.Run时nil问题
	// 故儿需要将其重新设置
	g := config.Appletuan
	g.Sub = g
	g.SubName = "appletuan"
	return g
}

//Dovoid 由 *AbstractSupport.Run 执行
func (tuan AppleTuan) DoVoid() {
	tuan.Do(tuan.Cookies)
}

//Do 最终执行签到
func (AppleTuan) Do(config interface{}) {
	if cookies, ok := config.(map[string]string); ok {
		tuan := autoAppletuan.AppleTuan{Cookies: cookies}
		tuan.Do()
	}
}
