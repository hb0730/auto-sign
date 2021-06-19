package support

import (
	"github.com/hb0730/auto-sign/application"
	"github.com/hb0730/auto-sign/config"
	"github.com/mritd/logger"
)

type PagodaWxMini struct {
	Support
}

var pagodaWx = application.PagodaWxMini{}

func init() {
	logger.Info("[support wx-mini-pagoda] 开始注册 ....")
	p := PagodaWxMini{}
	p.ISupport = p
	p.Name = "微信小程序-百果园"
	Register("pagodaWxMini", p)
}

func (m PagodaWxMini) DoRun() error {
	logger.Info("[support wx-mini-pagoda] 开始签到 ")
	result := GetPagodaWxMiniYaml()
	pagodaWx.Url = result.Url
	pagodaWx.Headers = result.Headers
	return pagodaWx.Start()
}

func GetPagodaWxMiniYaml() PagodaWxMiniYamlJson {
	y := config.ReadYaml()
	var result PagodaWxMiniYamlJson
	_ = y.Unmarshal(GetPagodaWxMiniYamlKey(), &result)
	return result
}

func GetPagodaWxMiniYamlKey() string {
	return "pagodaWxMini"
}

type PagodaWxMiniYamlJson struct {
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}
