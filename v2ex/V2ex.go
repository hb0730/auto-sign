package v2ex

import (
	"auto-sign/request"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"log"
)

type V2ex struct {
	cookies request.Cookies
}

const INDEX = "https://www.v2ex.com"

func (v *V2ex) Do() {
	v.checkin()
}
func (v *V2ex) checkin() {
	page := rod.New().MustConnect().MustPage("")
	defer page.MustClose()
	_ = page.SetCookies(convertCookie(v.cookies))
	page = page.MustNavigate(INDEX)
	page.Race().ElementR("a", "领取今日的登录奖励").MustHandle(func(el *rod.Element) {
		el.MustClick()
		page.MustElementR("input", "领取 X 铜币").MustClick()
		page.MustElementR(".message", "已成功领取每日登录奖励")
		log.Println("签到成功")
	}).Element(`.balance_area`).MustHandle(func(el *rod.Element) {
		log.Println("已经签过到了")
	}).MustDo()
}
func convertCookie(cookies request.Cookies) []*proto.NetworkCookieParam {
	c := make([]*proto.NetworkCookieParam, 0)
	for k, v := range cookies {
		c = append(c, &proto.NetworkCookieParam{Name: k, Value: v, Domain: "www.v2ex.com", HTTPOnly: true})
	}
	return c
}
