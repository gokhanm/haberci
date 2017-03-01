package mail

import (
    "crypto/tls"
    "gopkg.in/gomail.v2"
    "haberci/utils"
)

func MailSend(subject string, recipients_to []string, recipients_bcc []string, text string) {

    mail_conf := toml.Mail()

    m := gomail.NewMessage()
    m.SetHeader("From", mail_conf.From)

    if len(recipients_to) > 0 {
        m.SetHeader("To", recipients_to...)
    }

    if len(recipients_bcc) > 0 {
        m.SetHeader("Bcc", recipients_bcc...)
    }

    m.SetHeader("Subject", subject)
    m.SetBody("text/html", text)

    d := gomail.NewPlainDialer(mail_conf.Server, mail_conf.Port, mail_conf.Username, mail_conf.Password)
    d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }

}

