package v2ex

import (
	"auto-sign/util"
	"github.com/go-rod/rod"
)

type V2ex struct {
	cookies util.Cookies
}

const INDEX = "https://www.v2ex.com"

func (v *V2ex) Do() {
	v.checkin()
}
func (v *V2ex) checkin() {
	if len(v.cookies) == 0 {
		util.Warn("cookie len ==0")
		return
	}
	browser := rod.New().MustConnect()
	browser.MustSetCookies(util.ConvertCookies(v.cookies, ".v2ex.com"))
	defer browser.MustClose()
	page := browser.MustPage("")
	// 来自https://github.com/go-rod/v2ex-example
	page = page.MustNavigate(INDEX)
	page.Race().ElementR("a", "领取今日的登录奖励").MustHandle(func(el *rod.Element) {
		el.MustClick()
		page.MustElementR("input", "领取 X 铜币").MustClick()
		page.MustElementR(".message", "已成功领取每日登录奖励")
		util.Info("签到成功")
	}).Element(`.balance_area`).MustHandle(func(el *rod.Element) {
		util.Info("已经签过到了")
	}).MustDo()
}
