package config

import (
	"auto-sign/util"
	"bytes"
	"go.uber.org/config"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// ReadFile 读取配置文件
func ReadFile() (io.Reader, error) {
	workPath, _ := os.Executable()
	filePath2 := path.Dir(workPath)
	filePath2 = filepath.Join(filePath2, "/config/application.yml")
	content, err := ioutil.ReadFile(filePath2)
	return bytes.NewBuffer(content), err
}

//ReadYaml 读取yaml文件
func ReadYaml() (*config.YAML, error) {
	reader, err := ReadFile()
	if err != nil {
		util.ErrorF("read support file error, %v \n", err)
		return nil, err
	}
	return config.NewYAML(config.Source(reader))
}
