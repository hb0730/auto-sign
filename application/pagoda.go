package application

import (
	"encoding/json"
	"fmt"
	"github.com/hb0730/auto-sign/utils"
	"github.com/hb0730/auto-sign/utils/request"
	"github.com/mritd/logger"
)

//PagodaWxMini 百果园 微信小程序签到
type PagodaWxMini struct {
	Url     string
	Headers map[string]string
}

func (p *PagodaWxMini) Start() error {
	if p.Url == "" {
		return utils.AutoSignError{
			Module:  "wx-mini-pagoda",
			Method:  "sign",
			Message: "Url is null",
		}
	}
	if len(p.Headers) == 0 {
		return utils.AutoSignError{
			Module:  "wx-mini-pagoda",
			Method:  "sign",
			Message: "Headers is null",
		}
	}
	return p.sign()
}
func (p *PagodaWxMini) sign() error {
	rq, err := request.CreateRequest(
		"GET",
		p.Url,
		"")
	if err != nil {
		return err
	}
	rq.AddHeaders(p.Headers)
	err = rq.Do()
	if err != nil {
		return err
	}
	bt, err := rq.GetBody()
	if err != nil {
		return err
	}
	var result PagodaResult
	_ = json.Unmarshal(bt, &result)
	if result.ErrorCode == 0 || result.ErrorCode == 35702 {
		logger.Infof("[pagoda-wx-mini] sign success: [%s]", result.MessageInfo)
	} else {
		logger.Warnf("[pagoda-wx-mini] sign failed: [%s]", result.ErrorMsg)
		return utils.AutoSignError{
			Module:  "wx-mini-pagoda",
			Method:  "sign",
			Message: fmt.Sprintf("wx-mini-pagoda sign failed,message: [%s]", result.ErrorMsg),
		}
	}
	return nil
}

type PagodaResult struct {
	ErrorCode   int    `json:"errorCode"`
	ErrorMsg    string `json:"errorMsg"`
	MessageInfo string `yaml:"messageInfo"`
}
