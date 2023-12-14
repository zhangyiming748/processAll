package alert

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"log/slog"
	"strings"
	"time"
)

type Info struct {
	Form     string   `json:"form"`
	To       []string `json:"to"`
	Subject  string   `json:"subject"`
	Text     string   `json:"text"`
	Image    string   `json:"image"`
	Host     string   `json:"host"`
	Port     int      `json:"port"`
	Username string   `json:"username"`
	Password string   `json:"password"`
}

func init() {
	initLocal()
}

func initLocal() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
}

func Send(info *Info) {
	defer func() {
		if err := recover(); err != nil {
			slog.Warn("发送邮件发生错误", slog.String("错误原文", fmt.Sprint(err)))
		}
	}()
	m := gomail.NewMessage()
	m.SetHeader("From", info.Form)
	m.SetHeader("To", info.To...)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", info.Subject)
	m.SetBody("text/html", info.Text)
	if info.Image != "" {
		m.Attach("/home/Alex/lolcat.jpg")
	}
	d := gomail.NewDialer(info.Host, info.Port, info.Username, info.Password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	slog.Info("发送邮件", slog.Any("内容", info))
}
func (i *Info) SetFrom(s string) {
	i.Form = s
}
func (i *Info) SetTo(s []string) {
	i.To = s
}
func (i *Info) SetSubject(s string) {
	i.Subject = s
}

func (i *Info) SetText(s string) {
	i.Text = s
}
func (i *Info) AppendText(s string) {
	i.Text = strings.Join([]string{i.Text, s}, "<br>")
}

func (i *Info) SetImage(s string) {
	i.Image = s
}
func (i *Info) SetHost(s string) {
	i.Host = s
}
func (i *Info) SetPort(n int) {
	i.Port = n
}
func (i *Info) SetUsername(s string) {
	i.Username = s
}
func (i *Info) SetPassword(s string) {
	i.Password = s
}
