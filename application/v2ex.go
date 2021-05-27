package application

import (
	"github.com/go-rod/rod"
	"github.com/hb0730/auto-sign/utils"
	"time"
)

//https://v2x.com

//V2ex 通过Cookie签到
type V2ex struct {
	Cookies utils.Cookies
}

//Start 开始
func (v V2ex) Start() error {
	utils.Info("v2ex checkin .....")
	if len(v.Cookies) == 0 {
		utils.Warn("cookie len ==0")
		return &utils.AutoSignError{
			Module:  "v2ex",
			Method:  "start",
			Message: "cookies is null",
		}
	}
	return v.doStart()
}

func (v V2ex) doStart() error {
	b := utils.CreateBrowser(false)
	defer b.MustClose()
	page := b.MustSetCookies(utils.ConvertRodCookies(v.Cookies, "www.v2ex.com")...).
		MustPage("https://www.v2ex.com").
		MustWaitLoad()

	defer page.MustClose()
	page.Timeout(30*time.Second).
		Race().
		ElementR("a", "领取今日的登录奖励").
		MustHandle(func(e *rod.Element) {
			e.MustClick()
			page.MustElementR("input", "领取 X 铜币").MustClick()
			page.MustElementR(".message", "已成功领取每日登录奖励")
			utils.Info("v2ex 签到成功")
		}).Element(`.balance_area`).MustHandle(func(el *rod.Element) {
		utils.Info("v2ex 已经签过到了")
	}).MustDo()
	return nil
}
