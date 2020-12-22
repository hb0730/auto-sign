package send

import (
	"net/smtp"
	"strings"
)

type Mail struct {
	username string
	password string
	host     string
}

func (mail *Mail) Send(to string, subject string, content string) error {
	auth := smtp.PlainAuth("", mail.username, mail.password, mail.host)
	tos := strings.Split(to, ";")
	return smtp.SendMail(mail.host, auth, mail.username, tos, []byte(content))
}

func sendMsg(content string, subject string, to string) []byte {
	//[]byte("to: " +to+"\r\n"+"Subject: "+subject+"\r\n"+"")
	strBuilder := strings.Builder{}
	strBuilder.WriteString("to: ")
	strBuilder.WriteString(to)
	strBuilder.WriteString("\r\n")
	strBuilder.WriteString("Subject: ")
	strBuilder.WriteString(subject)
	strBuilder.WriteString("\r\n")
	strBuilder.WriteString("")
	return []byte(strBuilder.String())
}
