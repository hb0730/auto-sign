package application

import (
	"github.com/go-rod/rod"
	"github.com/hb0730/auto-sign/utils"
	"github.com/mritd/logger"
	"time"
)

// https://appletuan.com

// AppleTuan 通过cookie进行签到
type AppleTuan struct {
	Cookies utils.Cookies
}

//Start 开始
func (t AppleTuan) Start() error {
	logger.Info("[AppleTuan] checkins ....")
	if len(t.Cookies) <= 0 {
		logger.Warn("[AppleTuan] Cookies is null")
		return &utils.AutoSignError{
			Module:  "appletuan",
			Method:  "apple tuan sign",
			Message: "Cookies is null",
		}
	}
	return t.doStart()
}

func (t AppleTuan) doStart() error {
	b := utils.CreateBrowser(true)
	defer b.MustClose()
	page := b.MustSetCookies(utils.ConvertRodCookies(t.Cookies, "appletuan.com")...).
		MustPage("https://appletuan.com/checkins").
		MustWaitLoad()
	defer page.MustClose()
	page.Timeout(30*time.Second).
		Race().
		ElementR(`a[href="/checkins/start"]`, `签到`).
		MustHandle(func(e *rod.Element) {
			e.MustClick()
			page.MustElementR("span", `今日已签到`)
			logger.Info("[AppleTuan] 今日签到成功")
		}).ElementR("span", `今日已签到`).MustHandle(func(c *rod.Element) {
		logger.Info("[AppleTuan] 今日已签到成功")
	}).MustDo()
	return nil
}
