package browser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"sync"
)

var browser *rod.Browser
var runningMu sync.Mutex

//Create 创建*rod.Browser
func Create() *rod.Browser {
	runningMu.Lock()
	defer runningMu.Unlock()
	url := launcher.New().MustLaunch()
	b := rod.New().ControlURL(url).MustConnect()
	browser = b
	return b
}

//GetBrowser 获取创建的*rod.Browser,
//如果没有创建会自动创建
func GetBrowser() *rod.Browser {
	if browser == nil {
		Create()
	}
	return browser
}

func NewBrowser(headless bool) *rod.Browser {
	url := launcher.New().Headless(headless).MustLaunch()
	return rod.New().ControlURL(url).MustConnect()
}

//Close 关闭*rod.Browser#MustClose()
func Close() {
	if browser != nil {
		browser.MustClose()
	}
}
