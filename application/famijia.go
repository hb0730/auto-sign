package application

import (
	"encoding/json"
	"fmt"
	"github.com/hb0730/auto-sign/utils"
	"github.com/hb0730/go-request"
	"github.com/mritd/logger"
)

// Famijia Fa米家签到
type Famijia struct {
	Headers map[string]string
}

func (f Famijia) Start() error {
	logger.Info("[Famijia] sign start ...")
	if len(f.Headers) == 0 {
		logger.Warn("[Famijia] params is null")
		return utils.AutoSignError{
			Module:  "Famijia",
			Method:  "Start",
			Message: "Famijia params is null",
		}
	}
	return f.doStart()
}
func (f Famijia) doStart() error {
	logger.Info("[Famijia] sign ....")
	rq, err := request.CreateRequest(
		"POST",
		"https://fmapp.chinafamilymart.com.cn/api/app/market/member/signin/sign",
		"")
	if err != nil {
		return err
	}
	header := request.ConvertHeader(nil, f.otherHeaders())
	header = request.ConvertHeader(header, f.Headers)
	rq.Header(header)
	err = rq.Do()
	if err != nil {
		return err
	}
	bt, err := rq.GetBody()
	if err != nil {
		return err
	}
	var result FamijiaResult
	err = json.Unmarshal(bt, &result)
	if err != nil {
		return err
	}
	if result.Code == "200" || result.Code == "3004000" {
		logger.Infof("[Famijia] sign success,message:【%s】", result.Message)
	} else {
		logger.Warnf("[Famijia] sign failed message:【%s】", result.Message)
		return &utils.AutoSignError{
			Module:  "Famijia",
			Method:  "sign",
			Message: fmt.Sprintf("Famijia sign failed message:【%s】", result.Message),
		}
	}
	return nil
}

func (f Famijia) otherHeaders() map[string]string {
	return map[string]string{
		"Host":            "fmapp.chinafamilymart.com.cn",
		"Content-Type":    "application/json",
		"Accept":          "*/*",
		"loginChannel":    "app",
		"os":              "ios",
		"Accept-Language": "zh-Hans;q=1.0",
		"Accept-Encoding": "br;q=1.0, gzip;q=0.9, deflate;q=0.8",
		"User-Agent":      "Fa",
		"Connection":      "keep-alive",
		"fmVersion":       "2.4.1",
	}
}

type FamijiaResult struct {
	Code    json.Number `json:"code"`
	Message string      `json:"message"`
}
