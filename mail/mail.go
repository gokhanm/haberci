package mail

import (
    "crypto/tls"
    "gopkg.in/gomail.v2"
    "haberci/utils"
)

func MailSend(text string) {

    mail_conf := toml.Mail()
    addresses := make([]string, len(mail_conf.To))

    m := gomail.NewMessage()
    m.SetHeader("From", mail_conf.From)

    for i, recipient := range mail_conf.To {
        addresses[i] = m.FormatAddress(recipient, "")
    }

    m.SetHeader("To", addresses...)
    m.SetHeader("Subject", mail_conf.Subject)
    m.SetBody("text/html", text)

    d := gomail.NewPlainDialer(mail_conf.Server, mail_conf.Port, mail_conf.Username, mail_conf.Password)
    d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }

}

