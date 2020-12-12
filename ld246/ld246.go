package ld246

import (
	"auto-sign/request"
	"auto-sign/util"
	"encoding/json"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"log"
	"net/http"
)

const LOGIN_URL = "https://ld246.com/api/v2/login"
const LOGOUT_URL = "https://ld246.com/api/v2/logout"
const LD_INDEX = "https://ld246.com/"
const CHECKIN = "https://ld246.com/activity/checkin"

//
type LD struct {
	Username string
	Password string
}

//
func (ld *LD) Do() {
	if ld.Username == "" {
		log.Println("username is null")
		return
	}
	if ld.Password == "" {
		log.Println("password is null")
		return
	}
	r := ld.Login()
	cookies := setCookie(r.Token)
	ld.Index(cookies)
	ld.Checkin(cookies)
	ld.Logout(cookies)
}

func (ld *LD) Login() LoginResult {
	var result LoginResult
	log.Println("login .....")
	params := make(map[string]string, 2)
	params["userName"] = ld.Username
	params["userPassword"] = util.GetMd5(ld.Password)
	requestBody, _ := json.Marshal(params)
	headers := http.Header{}
	headers.Set("Content-Type", "application/json;charset=UTF-8")
	r := request.Request{Method: "POST", Url: LOGIN_URL, Params: string(requestBody)}
	req := r.CreateRequest()
	req.Header = headers
	body, isSuccess := request.Req(req, nil)
	if isSuccess {
		log.Println("login success")
		_ = json.Unmarshal([]byte(body), &result)
		return result
	}
	log.Println("login failed")
	return result
}
func (*LD) Checkin(cookies request.Cookies) {
	if len(cookies) == 0 {
		log.Println("token is null")
		return
	}
	c := util.ConvertCookies(cookies, ".ld246.com")
	u := launcher.New().StartURL("about:blank").MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect()
	defer browser.MustClose()
	browser.MustSetCookies(c)
	page := browser.MustSetCookies(c).MustPage("")
	page = page.MustNavigate(CHECKIN).MustWaitLoad()
	page.Race().ElementR("a", `领取今日签到奖励`).MustHandle(func(e *rod.Element) {
		e.MustClick()
		html := e.MustElement(`.btn`).MustText()
		log.Println(html)
	}).Element(`.btn`).MustHandle(func(c *rod.Element) {
		html := c.MustText()
		log.Println(html)
	}).MustDo()

}
func check(url string, cookies request.Cookies) {
	if len(cookies) == 0 || url == "" {
		log.Println("token is null")
		return
	}
	rq := request.Request{Method: "GET", Url: url, Params: ""}
	req := rq.CreateRequest()
	req.Header = setHeader()
	body, is := request.Req(req, cookies)
	if is {
		log.Printf("check success %v\n", body)
	}
}
func (*LD) Index(cookies request.Cookies) {
	if len(cookies) == 0 {
		log.Printf("token is null")
		return
	}
	r := request.Request{Method: "GET", Url: LD_INDEX, Params: ""}
	req := r.CreateRequest()
	req.Header = setHeader()
	body, is := request.Req(req, cookies)
	if is {
		log.Printf("request success %v\n", body)
		return
	}
	log.Printf("request index failed %v\n", body)

}

func (*LD) Logout(cookies request.Cookies) {
	if len(cookies) == 0 {
		log.Printf("token is null")
		return
	}
	r := request.Request{Method: "POST", Url: LOGOUT_URL, Params: ""}
	req := r.CreateRequest()
	req.Header = setHeader()
	body, is := request.Req(req, cookies)
	if is {
		log.Printf("logout success %v\n", body)
		return
	}
	log.Printf("logout failed %v\n", body)
}
func setCookie(token string) map[string]string {
	cookie := make(map[string]string, 0)
	cookie["symphony"] = token
	return cookie
}
func setHeader() http.Header {
	headers := http.Header{}
	headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36")
	headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	return headers
}

type LoginResult struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Token    string `json:"token"`
	UserName string `json:"username"`
}
