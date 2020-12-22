package config

import v2ex2 "auto-sign/v2ex"

type V2ex struct {
	AutoSign
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

func (v2 V2ex) GetConfig(config AutoSignConfig, typeInt int) interface{} {
	if typeInt == V2EX {
		return config.V2ex
	}
	return nil
}
func (v2 V2ex) Supports(config AutoSignConfig) Support {
	c := config.V2ex
	c.Sub = c
	c.SubName = "v2ex"
	return c
}