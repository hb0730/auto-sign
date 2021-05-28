package config

import (
	"github.com/hb0730/auto-sign/utils"
	"github.com/mritd/logger"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
)

var yaml *viper.Viper

// ReadYaml 读取 *viper.Viper
func ReadYaml() *viper.Viper {
	return yaml
}

// LoadYaml 重新加载配置并获取 *viper.Viper
func LoadYaml() *viper.Viper {
	initViper()
	return yaml
}

func initViper() {
	logger.Info(" [config] read yaml file init ...")
	workPath, _ := os.Executable()
	filePath := path.Dir(workPath)
	filePath = filepath.Join(filePath, "/config/application.yml")
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(&utils.AutoSignError{
			Module: "yaml",
			Method: "read yaml",
			E:      err,
		})
	}

	yaml = viper.GetViper()
}

func init() {
	initViper()
}
