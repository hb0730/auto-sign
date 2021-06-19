package application

import "testing"

func TestPagodaWxMini_Start(t *testing.T) {
	headers = map[string]string{
		"userToken":         "",
		"content-type":      "",
		"x-defined-verinfo": "",
	}
	pagoda := PagodaWxMini{
		Url:     "",
		Headers: headers,
	}
	_ = pagoda.Start()
}
