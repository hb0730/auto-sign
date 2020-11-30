package geekhub

import (
	"auto-sign/request"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
)

const URL = "https://geekhub.com/checkins/start"

const URL2 = "https://geekhub.com/checkins"

const authenticity_token = `<meta name="csrf-token" content="(.*?)" />`

var sessionId string

type Geekhub struct {
	SessionId string
}

func (geekhub *Geekhub) Do() {
	if geekhub.SessionId == "" {
		fmt.Println("session is  null")
		return
	}
	token := checkins(geekhub.SessionId)
	if token != "" {
		start(token)
	}
}

// checkins
func checkins(session string) string {
	fmt.Println("get token ...")
	headers := setSession(session)
	body, is := query("GET", URL2, "", headers)
	if !is {
		fmt.Println("session timout,reset session")
		return ""
	}
	compile := regexp.MustCompile(authenticity_token)
	token := compile.FindAllStringSubmatch(body, -1)
	if len(token) > 0 {
		t := token[0][1]
		fmt.Printf("token %s \n", t)
		return t
	}
	fmt.Println("获取token失败,签到失败")
	return ""
}

//start
func start(token string) {
	fmt.Println("start sign ....")
	values := url.Values{}
	// 转义
	values.Add("_method", "post")
	values.Add("authenticity_token", token)
	params := values.Encode()
	headers := setSession(sessionId)
	body, r := query("POST", URL, params, headers)
	if r {
		fmt.Printf("签到成功 %s\n", body)
		return
	}
	fmt.Printf("签到失败 %v\n", body)
}
func setSession(session string) http.Header {
	headers := http.Header{}
	headers.Set("Content-Type", "application/x-www-form-urlencoded")
	headers.Add("Cookie", fmt.Sprintf("_session_id=%s", session))
	return headers
}
func query(method string, url string, params string, header http.Header) (string, bool) {
	r := request.Request{Method: method, Url: url, Params: params, Headers: header}
	body, cookie, is := r.Request()
	if is {
		for _, v := range cookie {
			if v.Name == "_session_id" {
				sessionId = v.Value
			}
		}
		return body, true

	}
	return "", false
}
