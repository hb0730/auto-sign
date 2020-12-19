package config

import (
	"auto-sign/geekhub"
	"auto-sign/ld246"
	"auto-sign/util"
	v2ex2 "auto-sign/v2ex"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	GEEKHUB = iota
	LD246
	V2EX
)

var supportList = []Support{
	Geekhub{},
	Ld{},
	V2ex{},
}

type Geekhub struct {
	Cookies map[string]string `yaml:"cookies,omitempty"`
}

func (Geekhub) Support(config AutoSignConfig) Support {
	return config.Geekhub
}
func (g Geekhub) GetConfig(config AutoSignConfig, typeInt int) interface{} {
	if typeInt == GEEKHUB {
		return config.Geekhub
	}
	return nil
}
func (g Geekhub) Run() {
	g.Do(g.Cookies)
}
func (g Geekhub) Do(config interface{}) {
	if cookies, ok := config.(map[string]string); ok {
		hub := geekhub.Geekhub{Cookies: cookies}
		hub.Do()
	}
}

type Ld struct {
	User map[string]string `yaml:"user,omitempty"`
}

func (ld Ld) Do(config interface{}) {
	if user, ok := config.(map[string]string); ok {
		l := ld246.LD{Username: user["userName"], Password: user["password"]}
		l.Do()
	}
}
func (ld Ld) Run() {
	ld.Do(ld.User)
}
func (ld Ld) GetConfig(config AutoSignConfig, typeInt int) interface{} {
	if typeInt == LD246 {
		return config.Ld
	}
	return nil
}
func (ld Ld) Support(config AutoSignConfig) Support {
	return config.Ld
}

type V2ex struct {
	Cookies map[string]string `yaml:"cookies,omitempty"`
}

func (v2ex V2ex) Do(config interface{}) {
	if cookies, ok := config.(map[string]string); ok {
		v2 := v2ex2.V2ex{Cookies: cookies}
		v2.Do()
	}
}
func (v2 V2ex) Run() {
	v2.Do(v2.Cookies)
}
func (v2 V2ex) GetConfig(config AutoSignConfig, typeInt int) interface{} {
	if typeInt == V2EX {
		return config.V2ex
	}
	return nil
}
func (v2 V2ex) Support(config AutoSignConfig) Support {
	return config.V2ex
}

type AutoSignConfig struct {
	Geekhub Geekhub           `yaml:"geekhub"`
	Ld      Ld                `yaml:"ld246"`
	V2ex    V2ex              `yaml:"v2ex"`
	Cron    map[string]string `yaml:"cron"`
}

func (g AutoSignConfig) GetConfig(typeInt int) interface{} {
	for _, v := range supportList {
		result := v.GetConfig(g, typeInt)
		if result != nil {
			return result
		}
	}

	return nil
}

func Read() map[string]interface{} {
	content, err := ioutil.ReadFile("../config/application.yml")
	if err != nil {
		util.ErrorF("read yml file error, %v \n", err)
		return nil
	}
	rMap := make(map[string]interface{})
	err = yaml.Unmarshal(content, &rMap)
	if err != nil {
		util.ErrorF("read yml file error, %v \n", err)
		return nil
	}
	return rMap
}

type AutoSignError struct {
	Errors []string
}

func (err *AutoSignError) Error() string {
	return fmt.Sprintf("%v \n", err.Errors)
}

func RedStruct() (AutoSignConfig, error) {
	content, err := ioutil.ReadFile("config/application.yml")
	if err != nil {
		util.ErrorF("read yml file error, %v \n", err)
		return AutoSignConfig{}, &AutoSignError{Errors: []string{"读取yaml文件失败失败:message", err.Error()}}
	}
	var autoSign AutoSignConfig
	err = yaml.Unmarshal(content, &autoSign)
	if err != nil {
		util.ErrorF("read yml file error, %v \n", err)
		return AutoSignConfig{}, &AutoSignError{Errors: []string{"格式转换错误:message", err.Error()}}
	}
	return autoSign, nil
}
