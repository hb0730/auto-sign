package application

import "testing"

func TestFamijia_Start(t *testing.T) {
	f := Famijia{}
	f.Token = ""
	f.BlackBox = ""
	f.DeviceId = ""
	f.Start()
}
