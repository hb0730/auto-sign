package application

import "testing"

func TestFamijia_Start(t *testing.T) {
	headers := map[string]string{
		"token":    "",
		"blackBox": "",
		"deviceId": "",
	}
	f := Famijia{Headers: headers}
	f.Start()
}
