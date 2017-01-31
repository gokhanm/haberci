package mail

import (
    "crypto/tls"
    "gopkg.in/gomail.v2"
)

type Info struct {
    from string
    to string
    subject string
    smtp string
    password string
    port int
    username string
}


func MailSend(text string) {
    i := Info{}


    m := gomail.NewMessage()
    m.SetHeader("From", i.from)
    m.SetHeader("To", i.to)
    m.SetHeader("Subject", i.subject)
    m.SetBody("text", text)

    d := gomail.NewPlainDialer(i.smtp, i.port, i.username, i.password)
    d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }
}

