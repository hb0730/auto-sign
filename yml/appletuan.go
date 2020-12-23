package config

import autoAppletuan "auto-sign/appletuan"

// AppleTuan  https://appletuan.com/
type AppleTuan struct {
	AutoSign
	//Cookies 用于签到
	Cookies map[string]string `yaml:"cookies,omitempty"`
}

//Supports 所支持的，返回具体类型
func (AppleTuan) Supports(config AutoSignConfig) Support {
	g := config.Appletuan
	g.Sub = g
	g.SubName = "appletuan"
	return g
}

//Dovoid 有*AutoSign.Run 执行
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
