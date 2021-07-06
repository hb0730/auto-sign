package config

import (
	"github.com/hb0730/auto-sign/utils"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/mritd/logger"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
)

var k *koanf.Koanf
var ConfigPath string

func SetConfigPath(configPath string) {
	ConfigPath = configPath
}

func ReadYaml() *koanf.Koanf {
	if k == nil {
		return LoadKoanf("")
	}
	return k
}
func LoadKoanf(config string) *koanf.Koanf {
	load(config)
	return k
}
func load(config string) {
	k = koanf.New(".")
	if config == "" {
		defaultConfig()
	} else {
		err := k.Load(file.Provider(config), yaml.Parser())
		if err != nil {
			panic(&utils.AutoSignError{
				Module: "yaml",
				Method: "read yaml",
				E:      err,
			})
		}

	}
}

var v *viper.Viper

func GetViper() *viper.Viper {
	if v == nil {
		return LoadViper("")
	}
	return v
}

func LoadViper(config string) *viper.Viper {
	initViper(config)
	return v
}

func initViper(config string) {
	logger.Info("[config] read yaml file init ...")
	if config == "" {
		defaultConfig()
	} else {
		viper.SetConfigFile(config)
	}
	err := viper.ReadInConfig()
	if err != nil {
		panic(&utils.AutoSignError{
			Module: "yaml",
			Method: "read yaml",
			E:      err,
		})
	}
	v = viper.GetViper()
}

func defaultConfig() {
	workPath, _ := os.Executable()
	filePath := path.Dir(workPath)
	filePath = filepath.Join(filePath, "/config/application.yml")
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath(filePath)

	_ = k.Load(file.Provider("./config/application.yml"), yaml.Parser())
	_ = k.Load(file.Provider("../config/application.yml"), yaml.Parser())
	workPath, _ = os.Executable()
	filePath = path.Dir(workPath)
	filePath = filepath.Join(filePath, "/config/application.yml")
	_ = k.Load(file.Provider(filePath), yaml.Parser())
}
