package application

import (
	"encoding/json"
	"github.com/hb0730/auto-sign/utils"
	"github.com/mritd/logger"
)

//几鸡 https://cc.ax/

type ChainG struct {
	Cookies utils.Cookies
}

func (g ChainG) DoRun() error {
	if len(g.Cookies) == 0 {
		logger.Warn("[ChinaG] cookie size 0")
		return utils.AutoSignError{
			Module:  "ChinaG",
			Method:  "DoRun",
			Message: "Cookies is null",
		}
	}
	return g.doStart()
}

func (g ChainG) doStart() error {
	logger.Info("[ChinaG] sign start")

	req := utils.Request{
		Method: "POST",
		Url:    "https://cc.ax/user/checkin",
		Params: "",
	}
	request, err := req.CreateRequest()
	if err != nil {
		return err
	}
	response, err := utils.HttpRequest(request, g.Cookies)
	if err != nil {
		return err
	}
	bt, err := utils.GetBody(response)
	if err != nil {
		return err
	}
	var result GResult
	err = json.Unmarshal(bt, &result)
	if err != nil {
		logger.Error("[ChinaG] json反序列化错误")
		return utils.AutoSignError{
			Method:  "doStart",
			Module:  "ChinaG",
			Message: "json反序列化错误",
			E:       err,
		}
	}
	if result.Ret == 200 || result.Ret == 500 {
		logger.Info("[ChinaG] sign success")
		return nil
	}
	logger.Warnf("[ChinaG] sign failed,message: 【%s】", result.Msg)
	return utils.AutoSignError{
		Method:  "doStart",
		Module:  "ChinaG",
		Message: "签到失败,message: 【" + result.Msg + "】",
	}

}

type GResult struct {
	Ret int    `json:"ret"`
	Msg string `json:"msg"`
}
