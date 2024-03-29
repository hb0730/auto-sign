package support

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/hb0730/auto-sign/message"
	"github.com/mritd/logger"
	"time"
)

//支持的类型
var Supports = make(map[string]AutoRun, 0)

// ISupport 实际签到
type ISupport interface {
	DoRun() error
}

//Support 支持的类型
// 承上启下作用
type Support struct {
	Name string
	// AutoRun Cron调用
	AutoRun
	// ISupport 后置
	// 需要注册
	// hub := Geekhub{}
	//	hub.ISupport = hub
	ISupport
}

// Run Cron执行
func (s Support) Run() {
	logger.Info("cron 开始执行")
	err := rod.Try(func() {
		retry(s, 3)
	})
	if err != nil {
		logger.Error(err.Error())
		sendMessageError(s.Name, err)
	} else {
		sendSuccess(s.Name)
	}
}

// Register 将支持的类型进行注册
func Register(cronName string, support AutoRun) {
	Supports[cronName] = support
}

// retry 尝试机制
func retry(a Support, num int) {
	err := rod.Try(func() {
		err := a.DoRun()
		if err != nil {
			panic(err)
		}
	})
	if err != nil && num > 0 {
		for {
			num--
			time.Sleep(time.Duration(3) * time.Second)
			retry(a, num)
		}

	} else if err != nil {
		panic(err)
	}
}

// sendMessageError 发送错误信息
func sendMessageError(name string, err error) {
	var body = message.MessageBody{}
	body.Title = "签到失败"
	body.Content = fmt.Sprintf("%s,签到失败，错误信息: 【%s】", name, err.Error())
	send(body)
}

// sendSuccess 签到成功
func sendSuccess(name string) {
	body := message.MessageBody{
		Title:   "签到成功",
		Content: name + ",签到成功",
	}
	send(body)
}

// send 发送
func send(body message.MessageBody) {
	m := message.GetSupport()
	if m == nil {
		return
	}
	enabled := message.Enabled()
	if enabled {
		m.Send(body)
	}
}
