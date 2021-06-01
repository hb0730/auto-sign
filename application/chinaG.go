package application

import (
	"github.com/go-rod/rod"
	"github.com/hb0730/auto-sign/utils"
	"github.com/mritd/logger"
)

//几鸡 https://cc.ax/

type ChinaG struct {
	Username string
	Password string
}

func (g ChinaG) Start() error {
	if g.Username == "" || g.Password == "" {
		logger.Warn("[ChinaG] username/password is null")
		return utils.AutoSignError{
			Module:  "ChinaG",
			Method:  "DoRun",
			Message: "username/password is null",
		}
	}
	return g.doStart()
}

func (g ChinaG) doStart() error {
	b := utils.CreateBrowser(true)
	defer b.MustClose()
	//login
	page := b.MustPage("https://cc.ax/signin").MustWaitLoad()
	defer page.MustClose()

	page.
		MustElement(`input[name="email"]`).
		MustInput(g.Username)

	page.MustElement(`input[type="password"]`).
		MustInput(g.Password)

	page.MustElement(`.el-form-item__content > button`).MustClick().MustWaitLoad()
	page.MustElement(`.el-message-box > .el-message-box__btns > button`).MustClick()

	//sign
	//等待跳转
	page.MustWaitOpen()
	//等待页面渲染
	page.MustWaitLoad()
	// 关闭弹窗
	page.MustElement(".dialog-footer > button").MustClick().MustWaitLoad()

	page.Race().ElementR(`a`, "签到流量").MustHandle(func(e *rod.Element) {
		e.MustClick()
		logger.Info("[ChinaG] 签到成功")
	}).ElementR(`a`, "今日已签").MustHandle(func(e *rod.Element) {

		logger.Info("[ChinaG] 今日已签到")
	}).MustDo()

	//logout
	page.MustElementR(`a.nav-link[href="/user/logout" ]`, "登出").MustClick().MustWaitLoad()

	return nil
}
