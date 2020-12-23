package config

import "auto-sign/ld246"

type Ld struct {
	AutoSign
	User map[string]string `yaml:"user,omitempty"`
}

func (ld Ld) Do(config interface{}) {
	if user, ok := config.(map[string]string); ok {
		l := ld246.LD{Username: user["userName"], Password: user["password"]}
		l.Do()
	}
}

//func (ld Ld) Run() {
//	ld.Do(ld.User)
//}
func (ld Ld) DoVoid() {
	ld.Do(ld.User)
}

func (ld Ld) Supports(config AutoSignConfig) Support {
	c := config.Ld
	c.Sub = c
	c.SubName = "ld246"
	return c
}
