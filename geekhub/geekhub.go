package geekhub

import (
	"auto-sign/browser"
	"auto-sign/util"
	"github.com/go-rod/rod"
)

const GEEK_HUB = "https://www.geekhub.com/checkins"

type Geekhub struct {
	//SessionId string
	Cookies util.Cookies
}

func (geekhub *Geekhub) Do() {
	util.Info("geekhub checkin .....")

	if len(geekhub.Cookies) <= 0 {
		util.Warn("session is  null")
		return
	}
	geekhub.checkins()
}

// checkins
func (geekhub *Geekhub) checkins() {
	util.Info("get token ...")
	b := browser.NewBrowser(true)
	defer b.MustClose()
	page := b.MustSetCookies(util.ConvertCookies(geekhub.Cookies, "www.geekhub.com")).MustPage(GEEK_HUB).MustWaitLoad()
	page.Race().ElementR(`a[href="/checkins/start"]`, `签到`).MustHandle(func(e *rod.Element) {
		e.MustClick()
		page.MustElementR("span", `今日已签到`)
		util.Info("geekhub 今日签到成功")
	}).ElementR("span", `今日已签到`).MustHandle(func(c *rod.Element) {
		util.Info("geekhub 今日已签到成功")
	}).MustDo()
}
