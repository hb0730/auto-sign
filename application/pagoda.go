package application

import (
	"github.com/hb0730/auto-sign/utils"
	"github.com/hb0730/auto-sign/utils/request"
)

//PagodaWxMini 百果园 微信小程序签到
type PagodaWxMini struct {
	url     string
	headers map[string]string
}

func (p *PagodaWxMini) Start() error {
	if p.url == "" {
		return utils.AutoSignError{
			Module:  "wx-mini-pagoda",
			Method:  "sign",
			Message: "url is null",
		}
	}
	if len(p.headers) == 0 {
		return utils.AutoSignError{
			Module:  "wx-mini-pagoda",
			Method:  "sign",
			Message: "headers is null",
		}
	}
	return p.sign()
}
func (p *PagodaWxMini) sign() error {
	rq, err := request.CreateRequest(
		"GET",
		p.url,
		"")
	if err != nil {
		return err
	}
	rq.AddHeaders(p.headers)
	err = rq.Do()
	if err != nil {
		return err
	}
	bt, err := rq.GetBody()
	if err != nil {
		return err
	}
	result := string(bt)
	if result != "" {

	}
	return nil
}
