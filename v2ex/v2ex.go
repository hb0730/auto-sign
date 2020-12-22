package v2ex

import (
	"auto-sign/browser"
	"auto-sign/util"
	"github.com/go-rod/rod"
)

type V2ex struct {
	Cookies util.Cookies
}

const INDEX = "https://www.v2ex.com"

func (v *V2ex) Do() {
	util.Info("v2ex checkin .....")
	v.checkin()
}
func (v *V2ex) checkin() {
	if len(v.Cookies) == 0 {
		util.Warn("cookie len ==0")
		return
	}
	b := browser.NewBrowser(true)
	defer b.MustClose()
	// 来自https://github.com/go-rod/v2ex-example
	page := b.MustSetCookies(util.ConvertCookies(v.Cookies, "www.v2ex.com")).MustPage(INDEX).MustWaitLoad()
	//page = page.MustNavigate(INDEX)
	page.Race().ElementR("a", "领取今日的登录奖励").MustHandle(func(el *rod.Element) {
		el.MustClick()
		page.MustElementR("input", "领取 X 铜币").MustClick()
		page.MustElementR(".message", "已成功领取每日登录奖励")
		util.Info("v2ex 签到成功")
	}).Element(`.balance_area`).MustHandle(func(el *rod.Element) {
		util.Info("v2ex 已经签过到了")
	}).MustDo()
}
