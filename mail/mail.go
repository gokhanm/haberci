package mail

import (
    "crypto/tls"
    "gopkg.in/gomail.v2"
    "haberci/utils"
)

func MailSend(text string) {

    mail_conf := toml.Mail()

    m := gomail.NewMessage()
    m.SetHeader("From", mail_conf.From)

    m.SetHeader("To", mail_conf.To...)
    m.SetHeader("Subject", mail_conf.Subject)
    m.SetBody("text/html", text)

    d := gomail.NewPlainDialer(mail_conf.Server, mail_conf.Port, mail_conf.Username, mail_conf.Password)
    d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }

}

