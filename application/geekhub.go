package application

import (
	"github.com/go-rod/rod"
	"github.com/hb0730/auto-sign/utils"
	"time"
)

//  https://geekhub.com

// GeekHub 通过cookie进行签到
type GeekHub struct {
	//Cookies
	Cookies utils.Cookies
}

//Start 开始
func (g GeekHub) Start() error {
	utils.Info("geekhub checkin .....")
	if len(g.Cookies) <= 0 {
		utils.Warn("geekhub session is  null")
		return &utils.AutoSignError{
			Module:  "Geekhub",
			Method:  "Start",
			Message: "Geekhub Cookie is null",
		}
	}
	return g.doStart()
}

func (g GeekHub) doStart() error {
	b := utils.CreateBrowser(false)
	defer b.MustClose()
	page := b.MustSetCookies(utils.ConvertRodCookies(g.Cookies, "www.geekhub.com")...).
		MustPage("https://www.geekhub.com/checkins").
		MustWaitLoad()
	defer page.MustClose()
	page.Timeout(30*time.Second).
		Race().
		ElementR(`a[href="/checkins/start"]`, `签到`).
		MustHandle(func(e *rod.Element) {
			e.MustClick()
			page.MustElementR("span", `今日已签到`)
			utils.Info("geekhub 今日签到成功")
		}).ElementR("span", `今日已签到`).MustHandle(func(c *rod.Element) {
		utils.Info("geekhub 今日已签到成功")
	}).MustDo()

	return nil
}
