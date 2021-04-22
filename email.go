package email

import (
	"crypto/tls"
	"net"
	"net/smtp"
	"strconv"

	innerEmail "github.com/jordan-wright/email"
)

type Email struct {
	Username string
	Password string
	Host     string
	Port     int
}

func (e Email) Send(recipients []string, subject, body, file string) error {
	emailClient := innerEmail.NewEmail()
	emailClient.From = e.Username
	emailClient.To = recipients
	emailClient.Subject = subject
	emailClient.HTML = []byte(body)
	if file != "" {
		a, err := emailClient.AttachFile(file)
		if err != nil {
			return err
		}
		emailClient.Attachments = append(emailClient.Attachments, a)
	}
	addr := net.JoinHostPort(e.Host, strconv.Itoa(e.Port))
	auth := smtp.PlainAuth("", e.Username, e.Password, e.Host)
	if e.Port == 587 {
		return emailClient.Send(addr, auth)
	}
	return emailClient.SendWithTLS(addr, auth, &tls.Config{ServerName: e.Host})
}
