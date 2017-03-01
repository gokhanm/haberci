package mail

import (
    "crypto/tls"
    "gopkg.in/gomail.v2"
    "haberci/utils"
)

func MailSend(subject string, to_recipients []string, bcc_recipients []string, text string) {

    mail_conf := toml.Mail()

    m := gomail.NewMessage()
    m.SetHeader("From", mail_conf.From)

    if len(to_recipients) > 0 {
        m.SetHeader("To", to_recipients...)
    }

    if len(bcc_recipients) > 0 {
        m.SetHeader("Bcc", bcc_recipients...)
    }

    m.SetHeader("Subject", subject)
    m.SetBody("text/html", text)

    d := gomail.NewPlainDialer(mail_conf.Server, mail_conf.Port, mail_conf.Username, mail_conf.Password)
    d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }

}

