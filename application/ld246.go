package application

import (
	"encoding/json"
	"fmt"
	"github.com/go-rod/rod"
	"github.com/hb0730/auto-sign/utils"
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
		utils.Warn("username is null")
		return &utils.AutoSignError{
			Module:  "ld246",
			Method:  "start",
			Message: "username is null",
		}
	}
	if ld.Password == "" {
		utils.Warn("password is null")
		return &utils.AutoSignError{
			Module:  "ld246",
			Method:  "start",
			Message: "password is null",
		}
	}
	return ld.doStart()
}

func (ld Ld246) doStart() error {
	result := ld.Login()
	var cookies = map[string]string{
		"symphony": result.Token,
	}
	ld.Sign(cookies)

	ld.Logout(cookies)
	return nil
}

//Login 通过username/password 换取token
func (ld Ld246) Login() LoginResult {
	utils.Info("ld246 login .....")
	params := make(map[string]string, 0)
	params["userName"] = ld.Username
	params["userPassword"] = utils.GetMd5(ld.Password)
	requestBody, _ := json.Marshal(params)

	header := http.Header{}
	header.Set("Content-Type", "application/json;charset=UTF-8")

	req := utils.Request{
		Method: "POST",
		Url:    "https://ld246.com/api/v2/login",
		Params: string(requestBody),
	}
	request, err := req.CreateRequest()
	if err != nil {
		utils.Warn("create http request error")
		panic(&utils.AutoSignError{
			Module:  "ld246",
			Method:  "login",
			Message: "create http request error",
			E:       err,
		})
	}
	reponse, err := utils.HttpRequest(request, nil)
	if err != nil {
		utils.Warn("get response error")
		panic(&utils.AutoSignError{
			Module:  "ld246",
			Method:  "login",
			Message: "get response error",
			E:       err,
		})
	}
	by, err := utils.GetBody(reponse)
	if err != nil {
		utils.Warn("get response error")
		panic(&utils.AutoSignError{
			Module:  "ld246",
			Method:  "login",
			Message: "get response error",
			E:       err,
		})
	}
	utils.Info("login success")
	var result LoginResult
	_ = json.Unmarshal(by, &result)
	return result
}

// Sign 签到
func (ld Ld246) Sign(cookies utils.Cookies) {
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
			utils.Info(fmt.Sprintf("签到成功,%s \n", html))
		}).Element(`a.btn`).MustHandle(func(e *rod.Element) {
		html := e.MustText()
		str := fmt.Sprintf("今日已签到, %s \n", html)
		utils.Info(str)
	}).MustDo()
}

// Logout  登出
func (ld Ld246) Logout(cookies utils.Cookies) {
	if len(cookies) == 0 {
		utils.Warn("token is null")
		return
	}
	req := utils.Request{
		Method: "POST",
		Url:    "https://ld246.com/api/v2/logout",
		Params: "",
	}
	r, err := req.CreateRequest()
	if err != nil {
		utils.Warn("create http request error")
		panic(err)
	}
	r.Header = setHeader()
	reponse, err := utils.HttpRequest(r, cookies)
	if err != nil {
		utils.Warn("request error")
		panic(err)
	}
	by, err := utils.GetBody(reponse)
	if err != nil {
		utils.Warn("get body error")
		panic(err)
	}
	utils.InfoF("logout success %v\n", string(by))

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
