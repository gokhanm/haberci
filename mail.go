package main

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

// MailBody info
type MailBody struct {
	Subject       string
	ToRecipients  []string
	BccRecipients []string
	Message       string
}

// Send is send mail function
func (mb *MailBody) Send() error {

	m := gomail.NewMessage()
	m.SetHeader("From", Conf.Mail.From)

	if len(mb.ToRecipients) > 0 {
		m.SetHeader("To", mb.ToRecipients...)
	}

	if len(mb.BccRecipients) > 0 {
		m.SetHeader("Bcc", mb.BccRecipients...)
	}

	m.SetHeader("Subject", mb.Subject)
	m.SetBody("text/html", mb.Message)

	d := gomail.NewPlainDialer(Conf.Mail.Server, Conf.Mail.Port, Conf.Mail.Username, Conf.Mail.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return d.DialAndSend(m)
}
