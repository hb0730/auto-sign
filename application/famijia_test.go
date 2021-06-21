package application

import "testing"

func TestFamijia_Start(t *testing.T) {
	h := map[string]string{
		"token":    "",
		"blackBox": "",
		"deviceId": "",
	}
	f := Famijia{Headers: h}
	f.Start()
}
