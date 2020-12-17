package config

import (
	"auto-sign/geekhub"
	"auto-sign/ld246"
	"auto-sign/util"
	"auto-sign/v2ex"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Geekhub struct {
	Cookies map[string]string `yaml:"cookies,omitempty"`
}
type Ld struct {
	User map[string]string `yaml:"user,omitempty"`
}
type V2ex struct {
	Cookies map[string]string `yaml:"cookies,omitempty"`
}

func (Geekhub) Support(t interface{}) error {
	if _, ok := t.(Geekhub); ok {
		return nil
	}
	return &AutoSignError{Errors: []string{"类型不一致"}}
}
func (g Geekhub) Do(sign AutoSign) {
	hub := geekhub.Geekhub{Cookies: sign.Geekhub.Cookies}
	hub.Do()
}

func (V V2ex) Support(t interface{}) error {
	if _, ok := t.(V2ex); ok {
		return nil
	}
	return &AutoSignError{Errors: []string{"类型不一致"}}
}
func (V V2ex) Do(sign AutoSign) {
	v2 := v2ex.V2ex{Cookies: sign.V2ex.Cookies}
	v2.Do()
}

func (Ld) Support(t interface{}) error {
	if _, ok := t.(Ld); ok {
		return nil
	}
	return &AutoSignError{Errors: []string{"类型不一致"}}
}
func (l Ld) Do(sign AutoSign) {
	ld := ld246.LD{Username: sign.Ld.User["userName"], Password: sign.Ld.User["password"]}
	ld.Do()
}

type AutoSign struct {
	Geekhub Geekhub           `yaml:"geekhub"`
	Ld      Ld                `yaml:"ld246"`
	V2ex    V2ex              `yaml:"v2ex"`
	Cron    map[string]string `yaml:"cron"`
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

func RedStruct() (AutoSign, error) {
	content, err := ioutil.ReadFile("config/application.yml")
	if err != nil {
		util.ErrorF("read yml file error, %v \n", err)
		return AutoSign{}, &AutoSignError{Errors: []string{"读取yaml文件失败失败:message", err.Error()}}
	}
	var autoSign AutoSign
	err = yaml.Unmarshal(content, &autoSign)
	if err != nil {
		util.ErrorF("read yml file error, %v \n", err)
		return AutoSign{}, &AutoSignError{Errors: []string{"格式转换错误:message", err.Error()}}
	}
	return autoSign, nil
}
