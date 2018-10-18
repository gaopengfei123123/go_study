package mailer

import (
	"gopkg.in/gomail.v2"
)

// MailInfo 邮件代理信息
type MailInfo struct {
	Host       string
	Username   string
	Password   string
	Port       int
	Encryption string
}

// SendMail 发送邮件信息
type SendMail struct {
	CC, To, From []string
}

var (
	mailInfo MailInfo
	sendMail SendMail
)

func init() {
	mailInfo = MailInfo{
		Host:     "smtp.exmail.qq.com",
		Port:     465,
		Password: "xxx",
		Username: "gaopengfei@soyoung.com",
	}

	sendMail = SendMail{
		From: []string{"gaopengfei@soyoung.com"},
		CC:   []string{"gaopengfei@soyoung.com"},
		To:   []string{"5173180@qq.com"},
	}
}

// SendToMail 发送邮件
func SendToMail(subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", sendMail.From...)
	m.SetHeader("To", sendMail.To...)
	m.SetHeader("Cc", sendMail.CC...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(mailInfo.Host, mailInfo.Port, mailInfo.Username, mailInfo.Password)

	return d.DialAndSend(m)
}
