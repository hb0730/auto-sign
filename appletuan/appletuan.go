package appletuan

import (
	"auto-sign/browser"
	"auto-sign/util"
	"github.com/go-rod/rod"
)

const CHECKINS = "https://appletuan.com/checkins"

//AppleTuan 苹果团
type AppleTuan struct {
	Cookies util.Cookies
}

//Do 开始执行签到
func (tuan *AppleTuan) Do() {
	util.Info("AppleTuan checkins ....")
	if len(tuan.Cookies) == 0 {
		util.Warn("Cookies is null")
		return
	}
	tuan.checkins()
}

//checkins 执行签到
func (tuan *AppleTuan) checkins() {
	b := browser.NewBrowser(true)
	defer b.MustClose()
	page := b.MustSetCookies(util.ConvertCookies(tuan.Cookies, "appletuan.com")).MustPage(CHECKINS).MustWaitLoad()
	page.Race().ElementR(`a[href="/checkins/start"]`, `签到`).MustHandle(func(e *rod.Element) {
		e.MustClick()
		page.MustElementR("span", `今日已签到`)
		util.Info("AppleTuan 今日签到成功")
	}).ElementR("span", `今日已签到`).MustHandle(func(c *rod.Element) {
		util.Info("AppleTuan 今日已签到成功")
	}).MustDo()
}
