package v2ex

import (
	"auto-sign/request"
	"fmt"
	"net/url"
	"regexp"
)

type V2ex struct {
	cookies request.Cookies
}

const ONCE_REG = `once=(.*?)'`
const DAYILY_URL = "https://www.v2ex.com/mission/daily"
const STAR_URL = "https://www.v2ex.com/mission/daily/redeem"

func (v *V2ex) Do() {
	once := v.Dayily()
	v.Start(once)
}

func (v *V2ex) Start(once string) {
	if once == "" {
		fmt.Println("once is null")
		return
	}
	params := url.Values{}
	params.Set("once", once)
	body, is := request.Query("GET", STAR_URL, params.Encode(), v.cookies)
	if is {
		fmt.Println(body)
	}
}
func (v *V2ex) Dayily() string {
	if len(v.cookies) <= 0 {
		fmt.Println("cookie si null")
		return ""
	}
	body, is := request.Query("GET", DAYILY_URL, "", v.cookies)
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
