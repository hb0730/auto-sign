package ld246

import (
	"auto-sign/ld246"
	"auto-sign/yml"
)

type Ld struct {
	config.AbstractSupport
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

func (ld Ld) Supports(config config.YamlConfig) config.Support {
	// 这里的设置主要解决 *AbstractSupport.Run时nil问题
	// 故儿需要将其重新设置
	c := config.Ld
	c.Sub = c
	c.SubName = "ld246"
	return c
}
