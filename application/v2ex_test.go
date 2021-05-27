package application

import (
	"testing"
)

var cookies = map[string]string{
	"A2":          "",
	"PB3_SESSION": "",
}

func TestV2ex_Start(t *testing.T) {
	v := &V2ex{}
	v.Cookies = cookies
	v.Start()
}
