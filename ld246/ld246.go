package ld246

import (
	"auto-sign/request"
	"auto-sign/util"
	"encoding/json"
	"fmt"
	"net/http"
)

const LOGIN_URL = "https://ld246.com/api/v2/login"
const LOGOUT_URL = "https://ld246.com/api/v2/logout"
const LD_INDEX = "https://ld246.com/"

type LD struct {
	Username string
	Password string
}

func (ld *LD) Do() {
	if ld.Username == "" {
		fmt.Println("username is null")
		return
	}
	if ld.Password == "" {
		fmt.Println("password is null")
		return
	}
	r := ld.Login(ld.Username, ld.Password)
	ld.Index(r.Token)
	ld.Logout(r.Token)
}

func (*LD) Login(username string, password string) LoginResult {
	var r LoginResult
	fmt.Println("login .....")
	params := make(map[string]string, 2)
	params["userName"] = username
	params["userPassword"] = util.GetMd5(password)
	requestBody, _ := json.Marshal(params)
	headers := http.Header{}
	headers.Set("Content-Type", "application/json;charset=UTF-8")
	result, isSuccess := query("POST", LOGIN_URL, string(requestBody), headers)
	if isSuccess {
		fmt.Println("login success")
		_ = json.Unmarshal([]byte(result), &r)
		return r
	}
	fmt.Println("login failed")
	return r
}
func (*LD) Index(token string) {
	if token == "" {
		fmt.Printf("token is null")
		return
	}
	header := setCookie(token)
	body, is := query("GET", LD_INDEX, "", header)
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
	headers := setCookie(token)
	result, b := query("POST", LOGOUT_URL, "", headers)
	if b {
		fmt.Printf("logout success %v\n", result)
		return
	}
	fmt.Printf("logout failed %v\n", result)
}
func query(method string, url string, params string, header http.Header) (string, bool) {
	r := request.Request{Method: method, Url: url, Params: params, Headers: header}
	body, _, is := r.Request()
	if is {
		return body, true
	}
	return "", false
}

func setCookie(token string) http.Header {
	headers := http.Header{}
	if token == "" {
		return headers
	}
	headers.Add("Cookie", fmt.Sprintf("symphony=%s", token))
	return headers
}

type LoginResult struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Token    string `json:"token"`
	UserName string `json:"username"`
}
