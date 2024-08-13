package email

import (
	"github.com/go-gomail/gomail"
)

type SMTPEmailServiceInterface interface {
	ConfigAuth(user, pass string)
	ConfigServer(servidor string, port int)
	SendEmail(recipient, subject, body string) error
}

type EmailOptions struct {
	From    string
	Subject string
	Body    string
}

type SMTPGmailAdpter struct {
	user     string
	pass     string
	servidor string
	port     int
}

func (s *SMTPGmailAdpter) ConfigAuth(user, pass string) {
	s.user = user
	s.pass = pass
}

func (s *SMTPGmailAdpter) ConfigServer(servidor string, port int) {
	s.servidor = servidor
	s.port = port
}

func (s *SMTPGmailAdpter) SendEmail(recipient, subject, body string) error {
	d := gomail.NewDialer(s.servidor, s.port, s.user, s.pass)
	m := gomail.NewMessage()
	m.SetHeader("From", s.user)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	err := d.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
}
