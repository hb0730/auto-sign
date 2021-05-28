package support

import (
	"encoding/json"
	"github.com/hb0730/auto-sign/application"
	"github.com/hb0730/auto-sign/config"
	"github.com/hb0730/auto-sign/utils"
)

type Famijia struct {
	Support
}

func init() {
	utils.Info("[Famijia] 开始注册 ....")
	f := Famijia{}
	f.ISupport = f
	f.Name = "Fa米家"
	Register("famijia", f)
}

func (f Famijia) DoRun() error {
	utils.Info("[Famijia] 开始签到 ")
	yaml := config.ReadYaml()

	header := yaml.GetStringMap("famijia.headers")
	bt, _ := json.Marshal(header)
	var rest FaMiJia
	_ = json.Unmarshal(bt, &rest)

	fa := application.Famijia{}
	fa.Token = rest.Token
	fa.BlackBox = rest.BlackBox
	fa.DeviceId = rest.DeviceId
	return fa.Start()
}

type FaMiJia struct {
	Token    string `json:"token"`
	BlackBox string `json:"black_box"`
	DeviceId string `json:"device_id"`
}
