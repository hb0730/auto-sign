package support

import (
	config2 "auto-sign/config"
	"auto-sign/send"
	"auto-sign/util"
	"go.uber.org/config"
)

type Mail struct {
	Enabled  bool   `json:"enabled"`
	Host     string `json:"host"`
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	FromName string `json:"fromName"`
	To       string `json:"to"`
}

//Read 读取配置
func Read() (Mail, error) {
	reader, err := config2.ReadFile()
	if err != nil {
		util.ErrorF("read support file error, %v \n", err)
		return Mail{}, &ReadError{Errors: []string{"读取file失败:message", err.Error()}}
	}
	provider, err := config.NewYAML(config.Source(reader))
	if err != nil {
		util.ErrorF("read support file error, %v \n", err)
		return Mail{}, &ReadError{Errors: []string{"格式转换错误:message", err.Error()}}
	}
	var result Mail
	err = provider.Get("mail").Populate(&result)
	if err != nil {
		util.ErrorF("read support file error, %v \n", err)
		return Mail{}, &ReadError{Errors: []string{"获取mail配置失败:message", err.Error()}}
	}
	return result, nil
}

func (mail Mail) Send(subject string, content string) error {
	m := convert(mail)
	if m.Password == "" || m.Username == "" {
		return nil
	}
	return m.Send(subject, content, mail.To)
}

func (mail Mail) SendToArray(subject string, content string, to ...string) error {
	m := convert(mail)
	if m.Password == "" || m.Username == "" {
		return nil
	}
	return m.SendToArray(subject, content, to...)
}

func convert(m Mail) send.Mail {
	return send.Mail{
		Host:     m.Host,
		Protocol: m.Protocol,
		Port:     m.Port,
		Username: m.Username,
		Password: m.Password,
		FromName: m.FromName,
	}
}
