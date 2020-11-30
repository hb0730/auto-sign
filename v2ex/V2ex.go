package v2ex

import (
	"auto-sign/request"
	"fmt"
	"net/http"
	"regexp"
)

type V2ex struct {
	Cookie string
}

const ONCE_REG = `once=(.*?)'`
const DAYILY_URL = "https://www.v2ex.com/mission/daily"
const STAR_URL = "https://www.v2ex.com/mission/daily/redeem?once=%s"

func (v *V2ex) Do() {
	once := v.Dayily()
	v.Start(once)
}

func (v *V2ex) Start(once string) {
	if once == "" {
		fmt.Println("once is null")
		return
	}
	url := fmt.Sprintf(STAR_URL, once)
	headers := setCookie(v.Cookie)
	body, is := query("GET", url, "", headers)
	if is {
		fmt.Println(body)
	}
}
func (v *V2ex) Dayily() string {
	if v.Cookie == "" {
		fmt.Println("cookie si null")
		return ""
	}
	headers := setCookie(v.Cookie)
	body, is := query("GET", DAYILY_URL, "", headers)
	if is {
		compile := regexp.MustCompile(ONCE_REG)
		once := compile.FindAllStringSubmatch(body, -1)
		if len(once) > 0 {
			return once[0][1]
		}
		return ""
	}
	return ""
}
func query(method string, url string, params string, header http.Header) (string, bool) {
	r := request.Request{Method: method, Url: url, Params: params, Headers: header}
	body, _, is := r.Request()
	if is {
		return body, true

	}
	return "", false
}

func setCookie(cookie string) http.Header {
	headers := http.Header{}
	c := http.Cookie{Name: "A2", Value: cookie}
	headers.Set("Cookie", c.String())
	return headers
}
