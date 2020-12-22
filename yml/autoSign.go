package config

import (
	"auto-sign/util"
	"fmt"
)

//AutoSign 主要是做一个Abstract类，用于切面
type AutoSign struct {
	Support
	// Sub 用于*Support.Run()调用,防止丢失当前类
	//
	// 用法示例:
	//g := Geekhub{}
	// g.Cookies = map[string]string{"test": "测"}
	// g.Sub = g
	Sub     Support
	SubName string
}

func (support AutoSign) Run() {
	defer func() {
		if r := recover(); r != nil {
			util.ErrorF("run error %v \n", r)
			var c string
			subject := fmt.Sprintf("%s 执行出错", support.SubName)
			mail := Config.Mail
			if t, ok := r.(error); ok {
				c = t.Error()
			}
			content := fmt.Sprintf("auto-sign在执行cron时出错，具体详情\n: %v\n", c)
			mail.Send(subject, content, mail.To)
			panic(r)
		}
	}()
	// 抓取异常发送email
	if support.Sub != nil {
		support.Sub.DoVoid()
	}
}
