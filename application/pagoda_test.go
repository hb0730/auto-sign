package application

import "testing"

func TestPagodaWxMini_Start(t *testing.T) {
	h := map[string]string{
		"userToken": "",
	}
	pagoda := PagodaWxMini{
		Url:     "",
		Headers: h,
	}
	_ = pagoda.Start()
}
