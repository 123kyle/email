package email

import "gopkg.in/gomail.v2"

type GoMail struct {
	Client
}

func (g GoMail) Send(recipients []string, subject, body, file string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", g.Username)
	m.SetHeader("To", recipients...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	if file != "" {
		m.Attach(file)
	}
	d := gomail.NewDialer(g.Host, g.Port, g.Username, g.Password)
	return d.DialAndSend(m)
}
