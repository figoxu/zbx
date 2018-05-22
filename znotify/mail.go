package znotify

import (
	"gopkg.in/gomail.v2"
	"github.com/quexer/utee"
)

type MailConf struct {
	Host     string
	Port     int
	Username string
	Password string
}

func SendMail(conf MailConf, title, content string, to ...string) {
	m := gomail.NewMessage()
	m.SetHeader("From", conf.Username)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", title)
	m.SetBody("text/html", content)
	d := gomail.NewDialer(conf.Host, conf.Port, conf.Username, conf.Password)
	utee.Chk(d.DialAndSend(m))
}
