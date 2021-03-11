package geekhub

import (
	"auto-sign/browser"
	error2 "auto-sign/error"
	"auto-sign/util"
	"github.com/go-rod/rod"
	"time"
)

const GEEK_HUB = "https://www.geekhub.com/checkins"

type Geekhub struct {
	//SessionId string
	Cookies util.Cookies
}

func (geekhub *Geekhub) Do() error {
	util.Info("geekhub checkin .....")
	if len(geekhub.Cookies) <= 0 {
		util.Warn("geekhub session is  null")
		return &error2.AutoSignError{Module: "geekhub", Message: "session is  null"}
	}
	geekhub.RodPage()
	return nil
}

func (geekhub *Geekhub) RodPage() {
	util.Info("get token ...")
	b := browser.NewBrowser(true)
	defer b.MustClose()
	page := b.MustSetCookies(util.ConvertCookies(geekhub.Cookies, "www.geekhub.com")).MustPage(GEEK_HUB).MustWaitLoad()
	util.Retry(page, geekhub, 3)
}

func (Geekhub *Geekhub) Checking(page *rod.Page) {
	page.Timeout(30*time.Second).Race().ElementR(`a[href="/checkins/start"]`, `签到`).MustHandle(func(e *rod.Element) {
		e.MustClick()
		page.MustElementR("span", `今日已签到`)
		util.Info("geekhub 今日签到成功")
	}).ElementR("span", `今日已签到`).MustHandle(func(c *rod.Element) {
		util.Info("geekhub 今日已签到成功")
	}).MustDo()
}
