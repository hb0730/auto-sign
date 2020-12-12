package v2ex

import (
	"auto-sign/request"
	"auto-sign/util"
	"github.com/go-rod/rod"
	"log"
)

type V2ex struct {
	cookies request.Cookies
}

const INDEX = "https://www.v2ex.com"

func (v *V2ex) Do() {
	v.checkin()
}
func (v *V2ex) checkin() {
	if len(v.cookies) == 0 {
		log.Println("cookie len ==0")
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
		log.Println("签到成功")
	}).Element(`.balance_area`).MustHandle(func(el *rod.Element) {
		log.Println("已经签过到了")
	}).MustDo()
}
