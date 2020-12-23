package config

import (
	"auto-sign/util"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//AutoSignConfig 所支持的类型
// geekhub,appletuan,ld246,v2ex
// 读取yml装配
type AutoSignConfig struct {
	Geekhub   Geekhub           `yaml:"geekhub"`
	Appletuan AppleTuan         `yaml:"appletuan"`
	Ld        Ld                `yaml:"ld246"`
	V2ex      V2ex              `yaml:"v2ex"`
	Cron      map[string]string `yaml:"cron"`
	Mail      Mail              `yaml:"mail"`
}

//SupportsMap 当前支持的类型
// key 对应yaml的cron key
var SupportsMap = map[string]interface{}{
	"geekhub":   Geekhub{},
	"appletuan": AppleTuan{},
	"ld246":     Ld{},
	"v2ex":      V2ex{},
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
var Config AutoSignConfig

//RedStruct 读取当前配置，并转成struct
func RedStruct() (AutoSignConfig, error) {
	content, err := ioutil.ReadFile("config/application.yml")
	if err != nil {
		util.ErrorF("read yml file error, %v \n", err)
		return AutoSignConfig{}, &ReadError{Errors: []string{"读取yaml文件失败失败:message", err.Error()}}
	}
	var autoSign AutoSignConfig
	err = yaml.Unmarshal(content, &autoSign)
	if err != nil {
		util.ErrorF("read yml file error, %v \n", err)
		return AutoSignConfig{}, &ReadError{Errors: []string{"格式转换错误:message", err.Error()}}
	}
	Config = autoSign
	return autoSign, nil
}
