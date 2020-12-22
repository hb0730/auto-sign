package send

import (
	"auto-sign/util"
	"fmt"
	mail "github.com/xhit/go-simple-mail/v2"
	"sync"
)

type Mail struct {
	sync.Mutex
	Host     string
	Protocol string
	Port     int
	Username string
	Password string
	FromName string
}

var server *mail.SMTPServer

// 创建新的*mail.SMTPServer 只会创建一次
func (em *Mail) NewServer() *mail.SMTPServer {
	em.Lock()
	defer em.Unlock()
	if server == nil {
		client := mail.NewSMTPClient()
		client.Host = em.Host
		client.Port = em.Port
		client.Username = em.Username
		client.Password = em.Password
		// 是否加密
		client.Encryption = mail.EncryptionSSL
		server = client
		return client
	}
	return server
}

//GetServer 获取*mail.SMTPServer，可能为nil
func GetServer() *mail.SMTPServer {
	return server
}

func (em *Mail) Send(subject string, content string, to string) error {
	smtpClient, err := em.NewServer().Connect()
	if err != nil {
		util.ErrorF("Expected nil, got %v connecting to client", err)
		return err
	}

	email := mail.NewMSG()
	email.SetFrom(setFrom(em.FromName, em.Username)).
		AddTo(to).
		SetSubject(subject).
		SetBody(mail.TextHTML, content)
	err = email.Send(smtpClient)
	defer smtpClient.Close()
	return err

}
func (em *Mail) SendToArray(subject string, content string, to ...string) error {
	smtpClient, err := em.NewServer().Connect()
	if err != nil {
		util.ErrorF("Expected nil, got %v connecting to client", err)
		return err
	}

	email := mail.NewMSG()
	email.SetFrom(setFrom(em.FromName, em.Username)).
		AddTo(to...).
		SetSubject(subject).
		SetBody(mail.TextHTML, content)
	err = email.Send(smtpClient)
	defer smtpClient.Close()
	return err
}

func setFrom(from string, username string) string {
	return fmt.Sprintf("From %s <%s>", from, username)
}
