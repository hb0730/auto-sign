package application

import (
	"encoding/json"
	"fmt"
	"github.com/go-rod/rod"
	"github.com/hb0730/auto-sign/utils"
	"github.com/hb0730/auto-sign/utils/request"
	"github.com/mritd/logger"
	"net/http"
	"time"
)

//https://ld246.com

//headers 设置请求头，防止限流拦截
var headers = map[string]string{
	"User-Agent": "auto-sign/1.0.4",
	"Accept":     "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
}

//Ld246 通过token签到
type Ld246 struct {
	Username string
	Password string
}

//Start 开始
func (ld Ld246) Start() error {
	if ld.Username == "" {
		logger.Info("[ld246] username is null")
		return &utils.AutoSignError{
			Module:  "ld246",
			Method:  "start",
			Message: "username is null",
		}
	}
	if ld.Password == "" {
		logger.Warn("[ld246]  password is null")
		return &utils.AutoSignError{
			Module:  "ld246",
			Method:  "start",
			Message: "password is null",
		}
	}
	return ld.doStart()
}

func (ld Ld246) doStart() error {
	result, err := ld.Login()
	if err != nil {
		return err
	}
	var cookies = map[string]string{
		"symphony": result.Token,
	}
	ld.Sign(cookies)

	return ld.Logout(cookies)
}

//Login 通过username/password 换取token
func (ld Ld246) Login() (LoginResult, error) {
	logger.Info("[ld246]  login .....")
	var result LoginResult

	params := make(map[string]string, 0)
	params["userName"] = ld.Username
	params["userPassword"] = utils.GetMd5(ld.Password)
	requestBody, _ := json.Marshal(params)
	rq, err := request.CreateRequest(
		"POST",
		"https://ld246.com/api/v2/login",
		string(requestBody))
	if err != nil {
		return result, err
	}
	err = rq.Do()
	if err != nil {
		return result, err
	}
	by, err := rq.GetBody()
	if err != nil {
		return result, err
	}
	logger.Info("[ld246] login success")

	_ = json.Unmarshal(by, &result)
	return result, nil
}

// Sign 签到
func (ld Ld246) Sign(cookies map[string]string) {
	b := utils.CreateBrowser(true)
	defer b.MustClose()
	page := b.MustSetCookies(utils.ConvertRodCookies(cookies, ".ld246.com")...).
		MustPage("")
	defer page.MustClose()
	//设置header
	page.MustSetExtraHeaders(rodHeader()...)

	page = page.MustNavigate("https://ld246.com/activity/checkin").MustWaitLoad()

	page.Timeout(30*time.Second).
		Race().
		ElementR(`div.module__body > a.btn`, `领取今日签到奖励`).
		MustHandle(func(e *rod.Element) {
			_ = e.MustClick().WaitLoad()
			page.MustNavigate("https://ld246.com/activity/checkin").MustWaitLoad()
			html := page.MustElement("a.btn").MustWaitLoad().MustText()
			logger.Infof("[ld246] %s", fmt.Sprintf("签到成功,%s \n", html))
		}).Element(`a.btn`).MustHandle(func(e *rod.Element) {
		html := e.MustText()
		str := fmt.Sprintf("今日已签到, %s \n", html)
		logger.Infof("[ld246] %s", str)
	}).MustDo()
}

// Logout  登出
func (ld Ld246) Logout(cookies map[string]string) error {
	if len(cookies) == 0 {
		logger.Warn("[ld246] token is null")
		return nil
	}
	rq, err := request.CreateRequest(
		"POST",
		"https://ld246.com/api/v2/logout",
		"",
	)
	if err != nil {
		return err
	}
	rq.Header(setHeader())
	rq.AddCookiesFromMap(cookies)
	err = rq.Do()
	if err != nil {
		return err
	}
	by, err := rq.GetBody()
	logger.Infof("[ld246] logout success %v\n", string(by))
	return nil
}

type LoginResult struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Token    string `json:"token"`
	UserName string `json:"username"`
}

func rodHeader() []string {
	header := make([]string, 0)
	for k, v := range headers {
		header = append(append(header, k), v)
	}
	return header
}

func setHeader() http.Header {
	header := http.Header{}
	for k, v := range headers {
		header.Set(k, v)
	}
	return header
}
