package web

import (
	"github.com/hb0730/auto-sign/config"
	"github.com/hb0730/auto-sign/support"
)

func Run(auto support.AutoRun) {
	config.LoadViper(config.ConfigPath)
	config.LoadKoanf(config.ConfigPath)
	go func() {
		auto.Run()
	}()
}
