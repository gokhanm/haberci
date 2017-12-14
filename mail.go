package main

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

// MailSend is send mail function
func MailSend(subject string, toRecipients []string, bccRecipients []string, text string) {

	m := gomail.NewMessage()
	m.SetHeader("From", Conf.Mail.From)

	if len(toRecipients) > 0 {
		m.SetHeader("To", toRecipients...)
	}

	if len(bccRecipients) > 0 {
		m.SetHeader("Bcc", bccRecipients...)
	}

	m.SetHeader("Subject", subject)
	m.SetBody("text/html", text)

	d := gomail.NewPlainDialer(Conf.Mail.Server, Conf.Mail.Port, Conf.Mail.Username, Conf.Mail.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
