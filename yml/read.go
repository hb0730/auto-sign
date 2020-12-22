package config

import (
	"auto-sign/util"
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

type AutoSignConfig struct {
	Geekhub Geekhub           `yaml:"geekhub"`
	Ld      Ld                `yaml:"ld246"`
	V2ex    V2ex              `yaml:"v2ex"`
	Cron    map[string]string `yaml:"cron"`
	mail    Mail              `yaml:"mail"`
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

var Config AutoSignConfig

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
	Config = autoSign
	return autoSign, nil
}
