package support

import "testing"

func TestChinaG_Run(t *testing.T) {
	g := ChinaG{}
	g.Name = "chinaG"
	g.ISupport = g
	g.Run()
}
