package config

import "auto-sign/send"

type Mail struct {
	Host     string `yaml:"host"`
	Protocol string `yaml:"protocol"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	FromName string `yaml:"fromName"`
	To       string `yaml:"to"`
}

func (mail Mail) Send(subject string, content string, to string) error {
	m := convert(mail)
	if m.Password == "" || m.Username == "" {
		return nil
	}
	return m.Send(subject, content, to)
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
