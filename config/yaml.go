package config

import (
	"github.com/hb0730/auto-sign/utils"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
)

func ReadYaml() (*viper.Viper, error) {
	utils.Info("read yaml file")
	workPath, _ := os.Executable()
	filePath := path.Dir(workPath)
	filePath = filepath.Join(filePath, "/config/application.yml")
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(filePath)
	if err := viper.ReadInConfig(); err != nil {
		utils.ErrorF("read support file error, %v \n", err)
		return nil, &utils.AutoSignError{
			Module: "yaml",
			Method: "read yaml",
			E:      err,
		}
	}
	return viper.GetViper(), nil
}
