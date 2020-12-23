package config

import "auto-sign/geekhub"

type Geekhub struct {
	AutoSign
	Cookies map[string]string `yaml:"cookies,omitempty"`
}

func (Geekhub) Supports(config AutoSignConfig) Support {
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
