package utils

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

//Rod utils
// https://github.com/go-rod/rod

// CreateBrowser create Browser
func CreateBrowser(headless bool) *rod.Browser {
	url := launcher.
		New().
		Headless(headless).
		MustLaunch()
	return rod.New().ControlURL(url).MustConnect()
}

// ConvertRodCookies 将Cookies转换成rod Cookies
func ConvertRodCookies(cookies map[string]string, domain string) []*proto.NetworkCookie {
	array := make([]*proto.NetworkCookie, 0)
	for k, v := range cookies {
		array = append(array, &proto.NetworkCookie{Name: k, Value: v, Domain: domain})
	}
	return array
}
