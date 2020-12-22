package ld246

import (
	"auto-sign/browser"
	"auto-sign/util"
	"encoding/json"
	"fmt"
	"github.com/go-rod/rod"
	"net/http"
)

const LOGIN = "https://ld246.com/api/v2/login"
const LOGOUT = "https://ld246.com/api/v2/logout"
const CHECKIN = "https://ld246.com/activity/checkin"

var headers = map[string]string{
	"User-Agent": "auto-sign/1.0.4",
	"Accept":     "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
}

//
type LD struct {
	Username string
	Password string
}

//
func (ld *LD) Do() {
	util.Info("ld246 checkin .....")

	if ld.Username == "" {
		util.Warn("username is null")
		return
	}
	if ld.Password == "" {
		util.Warn("password is null")
		return
	}
	r := ld.Login()
	cookies := setCookie(r.Token)
	ld.Checkin(cookies)
	ld.Logout(cookies)
}

func (ld *LD) Login() LoginResult {
	var result LoginResult
	util.Info("login .....")
	params := make(map[string]string, 2)
	params["userName"] = ld.Username
	params["userPassword"] = util.GetMd5(ld.Password)
	requestBody, _ := json.Marshal(params)
	header := http.Header{}
	header.Set("Content-Type", "application/json;charset=UTF-8")
	r := util.Request{Method: "POST", Url: LOGIN, Params: string(requestBody)}
	req := r.CreateRequest()
	req.Header = header
	body, isSuccess := util.Req(req, nil)
	if isSuccess {
		util.Info("login success")
		_ = json.Unmarshal([]byte(body), &result)
		return result
	}
	util.Warn("login failed")
	return result
}
func (*LD) Checkin(cookies util.Cookies) {
	if len(cookies) == 0 {
		util.Warn("token is null")
		return
	}
	b := browser.NewBrowser(true)
	defer b.MustClose()
	page := b.MustSetCookies(util.ConvertCookies(cookies, ".ld246.com")).MustPage("")
	page.MustSetExtraHeaders(convertHeader()...)
	page.MustNavigate(CHECKIN).MustWaitLoad()
	page.Race().ElementR(`a[class="btn green"]`, `领取今日签到奖励`).MustHandle(func(e *rod.Element) {
		e.MustClick()
		html := e.MustElement(`a[class="btn"]`).MustText()
		util.Info(fmt.Sprintf("签到成功, %s \n", html))
	}).Element(`a[class="btn"]`).MustHandle(func(e *rod.Element) {
		html := e.MustText()
		str := fmt.Sprintf("今日已签到, %s \n", html)
		util.Info(str)
	}).MustDo()
}

func (*LD) Logout(cookies util.Cookies) {
	if len(cookies) == 0 {
		util.Warn("token is null")
		return
	}
	r := util.Request{Method: "POST", Url: LOGOUT, Params: ""}
	req := r.CreateRequest()
	req.Header = setHeader()
	body, is := util.Req(req, cookies)
	if is {
		util.InfoF("logout success %v\n", body)
		return
	}
	util.WarnF("logout failed %v\n", body)
}
func setCookie(token string) map[string]string {
	if len(token) == 0 {
		return nil
	}
	cookie := make(map[string]string, 0)
	cookie["symphony"] = token
	return cookie
}
func setHeader() http.Header {
	header := http.Header{}
	for k, v := range headers {
		header.Set(k, v)
	}
	return header
}
func convertHeader() []string {
	header := make([]string, 0)
	for k, v := range headers {
		header = append(append(header, k), v)
	}
	return header
}

type LoginResult struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Token    string `json:"token"`
	UserName string `json:"username"`
}
