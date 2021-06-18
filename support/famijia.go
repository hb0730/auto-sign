package support

import (
	"github.com/hb0730/auto-sign/application"
	"github.com/hb0730/auto-sign/config"
	"github.com/mritd/logger"
)

type Famijia struct {
	Support
}

var fa = application.Famijia{}

func init() {
	logger.Info("[support famijia] 开始注册 ....")
	f := Famijia{}
	f.ISupport = f
	f.Name = "Fa米家"
	Register("famijia", f)
}

func (f Famijia) DoRun() error {
	logger.Info("[support famijia] 开始签到 ")
	k := config.ReadYaml()
	header := k.StringMap("famijia.headers")
	fa.Headers = header
	return fa.Start()
}
