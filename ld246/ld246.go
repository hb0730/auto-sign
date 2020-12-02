package ld246

import (
	"auto-sign/request"
	"auto-sign/util"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
)

//
//import (
//	"auto-sign/request"
//	"auto-sign/util"
//	"encoding/json"
//	"fmt"
//	"net/http"
//	"net/url"
//	"regexp"
//)
//
const LOGIN_URL = "https://ld246.com/api/v2/login"
const LOGOUT_URL = "https://ld246.com/api/v2/logout"
const LD_INDEX = "https://ld246.com/"
const CHECKIN = "https://ld246.com/activity/checkin"
const CHECK = "https://ld246.com/activity/daily-checkin"
const CSRFTOKEN_REG = `csrfToken: '(.*?)'`

//
type LD struct {
	Username string
	Password string
}

//
func (ld *LD) Do() {
	if ld.Username == "" {
		fmt.Println("username is null")
		return
	}
	if ld.Password == "" {
		fmt.Println("password is null")
		return
	}
	r := ld.Login()
	ld.Index(r.Token)
	ld.Checkin(r.Token)
	ld.Logout(r.Token)
}

func (ld *LD) Login() LoginResult {
	var result LoginResult
	fmt.Println("login .....")
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
		fmt.Println("login success")
		_ = json.Unmarshal([]byte(body), &result)
		return result
	}
	fmt.Println("login failed")
	return result
}

func (*LD) Checkin(token string) {
	if token == "" {
		fmt.Println("token is null")
		return
	}
	r := request.Request{Method: "GET", Url: CHECKIN, Params: ""}
	req := r.CreateRequest()
	req.Header = setHeader()
	body, is := request.Req(req, setCookie(token))
	if is {
		compile := regexp.MustCompile(CSRFTOKEN_REG)
		once := compile.FindAllStringSubmatch(body, -1)
		fmt.Println("check start ....")
		chek(token, once[0][1])
		return
	}
	fmt.Printf("request index failed %v\n", body)
}

func chek(token string, csrfToken string) {
	if csrfToken == "" {
		fmt.Println("csrfToken is null")
		return
	}
	params := url.Values{}
	params.Set("token", csrfToken)
	r := request.Request{Method: "GET", Url: CHECK, Params: params.Encode()}
	req := r.CreateRequest()
	req.Header = setHeader()
	body, is := request.Req(req, setCookie(token))
	if is {
		fmt.Printf("check success %v\n", body)
	}
}

func (*LD) Index(token string) {
	if token == "" {
		fmt.Printf("token is null")
		return
	}
	r := request.Request{Method: "GET", Url: LD_INDEX, Params: ""}
	req := r.CreateRequest()
	req.Header = setHeader()
	body, is := request.Req(req, setCookie(token))
	if is {
		fmt.Printf("request success %v\n", body)
		return
	}
	fmt.Printf("request index failed %v\n", body)

}

func (*LD) Logout(token string) {
	if token == "" {
		fmt.Printf("token is null")
		return
	}
	r := request.Request{Method: "POST", Url: LOGOUT_URL, Params: ""}
	req := r.CreateRequest()
	req.Header = setHeader()
	body, is := request.Req(req, setCookie(token))
	if is {
		fmt.Printf("logout success %v\n", body)
		return
	}
	fmt.Printf("logout failed %v\n", body)
}
func setCookie(token string) map[string]string {
	cookie := make(map[string]string, 0)
	cookie["symphony"] = token
	return cookie
}
func setHeader() http.Header {
	headers := http.Header{}
	headers.Set("User-Agent", "hb0730/1.0.0")
	return headers
}

type LoginResult struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Token    string `json:"token"`
	UserName string `json:"username"`
}
