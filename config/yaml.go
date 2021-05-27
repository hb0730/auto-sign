package config

import (
	"github.com/hb0730/auto-sign/utils"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
)

var Viper *viper.Viper

func ReadYaml() *viper.Viper {
	return Viper
}

func initViper() {
	utils.Info("read yaml file init ...")
	utils.Info("read yaml file")
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

	Viper = viper.GetViper()
}

func init() {
	initViper()
}
