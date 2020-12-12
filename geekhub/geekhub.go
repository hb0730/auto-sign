package geekhub

import (
	"auto-sign/request"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

const URL = "https://www.geekhub.com/checkins/start"

const URL2 = "https://www.geekhub.com/checkins"

const authenticity_token = `<meta name="csrf-token" content="(.*?)" />`

type Geekhub struct {
	//SessionId string
	Cookies request.Cookies
}

func (geekhub *Geekhub) Do() {
	if len(geekhub.Cookies) <= 0 {
		log.Println("session is  null")
		return
	}
	token := geekhub.checkins()
	if token != "" {
		geekhub.start(token)
	}
}

// checkins
func (geekhub *Geekhub) checkins() string {
	log.Println("get token ...")
	r := request.Request{Method: "GET", Url: URL2, Params: ""}
	req := r.CreateRequest()
	req.Header = setHeader()
	request.SetCookie(geekhub.Cookies, req)
	body, is := geekhub.do(req)
	if !is {
		log.Println("session timout,reset session")
		return ""
	}
	compile := regexp.MustCompile(authenticity_token)
	token := compile.FindAllStringSubmatch(body, -1)
	if len(token) > 0 {
		t := token[0][1]
		log.Printf("token %s \n", t)
		return t
	}
	log.Println("获取token失败,签到失败")
	return ""
}

//start
func (geekhub *Geekhub) start(token string) {
	log.Println("start sign ....")
	values := url.Values{}
	// 转义
	values.Add("_method", "post")
	values.Add("authenticity_token", token)
	params := values.Encode()
	r := request.Request{Method: "POST", Url: URL, Params: params}
	req := r.CreateRequest()
	req.Header = setHeader()
	request.SetCookie(geekhub.Cookies, req)
	body, is := geekhub.do(req)
	if is {
		log.Printf("签到成功 %s\n", body)
		return
	}
	log.Printf("签到失败 %v\n", body)
}
func (geekhub *Geekhub) do(req *http.Request) (string, bool) {
	body, cookie, is := request.ClientDo(req)
	if is {
		for _, v := range cookie {
			if v.Name == "_session_id" {
				geekhub.Cookies["_session_id"] = v.Value
			}
		}
		return body, is
	}
	return "", false
}
func setHeader() http.Header {
	headers := http.Header{}
	headers.Set("Content-Type", "application/x-www-form-urlencoded")
	return headers
}
