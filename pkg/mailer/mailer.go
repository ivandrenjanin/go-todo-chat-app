package mailer

import "gopkg.in/gomail.v2"

type Mailer struct {
	dialer *gomail.Dialer
}

func New(host, username, password string, port int) Mailer {
	d := gomail.NewDialer(host, port, username, password)

	return Mailer{
		dialer: d,
	}
}

func (m Mailer) Send(to, subject, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "no-reply@localhost.com")
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	if err := m.dialer.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
