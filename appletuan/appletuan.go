package appletuan

import (
	"auto-sign/browser"
	error2 "auto-sign/error"
	"auto-sign/util"
	"github.com/go-rod/rod"
	"time"
)

const CHECKINS = "https://appletuan.com/checkins"

//AppleTuan 苹果团
type AppleTuan struct {
	Cookies util.Cookies
}

func (tuan *AppleTuan) Do() error {
	util.Info("AppleTuan checkins ....")
	if len(tuan.Cookies) == 0 {
		util.Warn("Cookies is null")
		return &error2.AutoSignError{Module: "appleTuan", Message: "Cookies is null"}
	}
	tuan.RodPage()
	return nil
}

func (tuan *AppleTuan) RodPage() {
	b := browser.NewBrowser(true)
	defer b.MustClose()
	page := b.MustSetCookies(util.ConvertCookies(tuan.Cookies, "appletuan.com")).MustPage(CHECKINS).MustWaitLoad()
	util.Retry(page, tuan, 3)
}

func (tuan *AppleTuan) Checking(page *rod.Page) {
	page.Timeout(30*time.Second).Race().ElementR(`a[href="/checkins/start"]`, `签到`).MustHandle(func(e *rod.Element) {
		e.MustClick()
		page.MustElementR("span", `今日已签到`)
		util.Info("AppleTuan 今日签到成功")
	}).ElementR("span", `今日已签到`).MustHandle(func(c *rod.Element) {
		util.Info("AppleTuan 今日已签到成功")
	}).MustDo()
}
