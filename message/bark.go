package message

import (
	"encoding/json"
	"github.com/hb0730/auto-sign/config"
	"github.com/hb0730/auto-sign/utils/request"
	"github.com/mritd/logger"
)

// https://github.com/Finb/bark-server

// Bark 用于ios推送
type Bark struct {
	url string
	key string
}

var b = Bark{}

func init() {
	logger.Info("[message bark] start ...")
	yaml := config.ReadYaml()
	bark := yaml.StringMap("message.bark")
	b.url = bark["url"]
	b.key = bark["key"]
	Register("bark", b)
}

func (b Bark) Send(messageBody MessageBody) {
	var url = b.url
	body := requestBody{}
	body.DeviceKey = b.key
	body.Category = "auto-sign"
	body.Title = messageBody.Title
	body.Body = messageBody.Content
	bt, _ := json.Marshal(body)
	rq, err := request.CreateRequest(
		"POST",
		url,
		string(bt),
	)
	if err != nil {
		logger.Errorf("[message bark] 发送失败  error message 【%s】", err.Error())
		return
	}
	rq.AddHeader("Content-Type", "application/json; charset=utf-8")
	err = rq.Do()
	if err != nil {
		logger.Errorf("[message bark] 发送失败  error message 【%s】", err.Error())
	}
}

type requestBody struct {
	DeviceToken string            `json:"device_token"`
	DeviceKey   string            `json:"device_key"`
	Category    string            `json:"category"`
	Title       string            `json:"title"`
	Body        string            `json:"body"`
	Sound       string            `json:"sound"`
	ExtParams   map[string]string `json:"ext_params"`
}
