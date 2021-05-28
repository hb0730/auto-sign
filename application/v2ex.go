package application

import (
	"github.com/go-rod/rod"
	"github.com/hb0730/auto-sign/utils"
	"github.com/mritd/logger"
	"time"
)

//https://v2x.com

//V2ex 通过Cookie签到
type V2ex struct {
	Cookies utils.Cookies
}

//Start 开始
func (v V2ex) Start() error {
	logger.Info("[v2ex] checkin .....")
	if len(v.Cookies) == 0 {
		logger.Warn("[v2ex] cookie len ==0")
		return &utils.AutoSignError{
			Module:  "v2ex",
			Method:  "start",
			Message: "cookies is null",
		}
	}
	return v.doStart()
}

func (v V2ex) doStart() error {
	b := utils.CreateBrowser(true)
	defer b.MustClose()
	// 来自https://github.com/go-rod/v2ex-example
	page := b.MustSetCookies(utils.ConvertRodCookies(v.Cookies, ".v2ex.com")...).MustPage("")
	defer page.MustClose()
	page.MustSetExtraHeaders("accept-language", "zh-CN,zh;q=0.9")
	page.MustNavigate("https://www.v2ex.com/").MustWaitLoad()

	page.Timeout(30*time.Second).
		Race().
		ElementR("a", "领取今日的登录奖励").
		MustHandle(func(e *rod.Element) {
			e.MustClick()
			page.MustElementR("input", "领取 X 铜币").MustClick()
			page.MustElementR(".message", "已成功领取每日登录奖励")
			logger.Info("[v2ex] 签到成功")
		}).Element(`.balance_area`).MustHandle(func(el *rod.Element) {
		logger.Info("[v2ex] 已经签过到了")
	}).MustDo()
	return nil
}
