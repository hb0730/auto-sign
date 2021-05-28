package message

import (
	"encoding/json"
	"fmt"
	"github.com/hb0730/auto-sign/config"
	"github.com/hb0730/auto-sign/utils"
	mail2 "github.com/xhit/go-simple-mail/v2"
	"sync"
)

type Mail struct {
	Host     string `json:"host" yaml:"host"`
	Protocol string `json:"protocol" yaml:"protocol"`
	Port     int    `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	FromName string `json:"from_name" yaml:"from_name"`
	To       string `json:"to" yaml:"to"`
}

var mail Mail

func init() {
	utils.Info("mail start ...")
	yaml := config.ReadYaml()
	mailMap := yaml.GetStringMap("message.mail")
	bt, _ := json.Marshal(mailMap)
	_ = json.Unmarshal(bt, &mail)
	Register("mail", mail)
}

func (m Mail) Send(message MessageBody) {
	eml := NewServer()
	client, err := eml.Connect()
	if err != nil {
		utils.ErrorF("[mail] 发送失败 message error: 【%s】", err.Error())
		return
	}
	msg := mail2.NewMSG()
	msg.SetFrom(setFrom(m.FromName, m.Username)).
		AddTo(m.To).
		SetSubject(message.Title).
		SetBody(mail2.TextHTML, message.Content)
	_ = msg.Send(client)
	defer client.Close()
}

var server *mail2.SMTPServer
var mutex sync.Mutex

// NewServer 创建 SMTPServer 服务
func NewServer() *mail2.SMTPServer {
	mutex.Lock()
	defer mutex.Unlock()
	if server == nil {
		client := mail2.NewSMTPClient()
		client.Host = mail.Host
		client.Port = mail.Port
		client.Password = mail.Password
		client.Username = mail.Username
		client.Encryption = mail2.EncryptionSSL
		server = client
	}
	return server
}

func setFrom(from string, username string) string {
	return fmt.Sprintf("From %s <%s>", from, username)
}
