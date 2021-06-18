package support

import (
	"github.com/hb0730/auto-sign/config"
	"testing"
)

func TestFamijia_DoRun(t *testing.T) {
	config.LoadKoanf()
	f := Famijia{}
	f.Name = "famijia"
	f.ISupport = f
	f.DoRun()
}
