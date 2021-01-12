package config

import (
	"auto-sign/util"
	"auto-sign/yml/appletuan"
	"auto-sign/yml/geekhub"
	"auto-sign/yml/ld246"
	"auto-sign/yml/v2ex"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//YamlConfig 所支持的类型
// geekhub,appletuan,ld246,v2ex
// 读取yml装配
type YamlConfig struct {
	Geekhub   geekhub.Geekhub     `yaml:"geekhub"`
	Appletuan appletuan.AppleTuan `yaml:"appletuan"`
	Ld        ld246.Ld            `yaml:"ld246"`
	V2ex      v2ex.V2ex           `yaml:"v2ex"`
	Cron      map[string]string   `yaml:"cron"`
	Mail      Mail                `yaml:"mail"`
}

//SupportsMap 当前支持的类型
// key 对应yaml的cron key
var SupportsMap = map[string]interface{}{
	"geekhub":   geekhub.Geekhub{},
	"appletuan": appletuan.AppleTuan{},
	"ld246":     ld246.Ld{},
	"v2ex":      v2ex.V2ex{},
}

//Read 读取配置文件
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

//ReadError 读取错误
type ReadError struct {
	Errors []string
}

//Error 错误信息
func (err *ReadError) Error() string {
	return fmt.Sprintf("%v \n", err.Errors)
}

//Config 用于获取当前的email配置
var Config YamlConfig

//RedStruct 读取当前配置，并转成struct
func RedStruct() (YamlConfig, error) {
	content, err := ioutil.ReadFile("config/application.yml")
	if err != nil {
		util.ErrorF("read yml file error, %v \n", err)
		return YamlConfig{}, &ReadError{Errors: []string{"读取yaml文件失败失败:message", err.Error()}}
	}
	var autoSign YamlConfig
	err = yaml.Unmarshal(content, &autoSign)
	if err != nil {
		util.ErrorF("read yml file error, %v \n", err)
		return YamlConfig{}, &ReadError{Errors: []string{"格式转换错误:message", err.Error()}}
	}
	Config = autoSign
	return autoSign, nil
}
