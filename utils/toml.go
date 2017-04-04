package toml

import (
    "errors"

    "github.com/BurntSushi/toml"
)

var ConfigPath string

type MailConfig struct {
    Server string   `toml:"server"`
    Port int        `toml:"port"`
    Username string `toml:"username"`
    Password string `toml:"password"`
    From string     `toml:"from"`
}

type YtsConfig struct{
    Enabled string       `toml:enabled`
    PageLimit string    `toml:"pageLimit"`
    ToRecipients []string `toml:"to_recipients"`
    BccRecipients []string `toml:"bcc_recipients"`
    Subject string      `toml:"subject"`
}

type tConf struct {
    Title string
    Mail MailConfig
    Yts YtsConfig
}

func Load(conf string) {
    ConfigPath = conf
}

func Parse() (*tConf, error) {
    var conf *tConf

    if _, err := toml.DecodeFile(ConfigPath, &conf); err != nil {
        panic(errors.New("haberci.toml conf file not found."))
    }

    return conf, nil
}

func Mail() MailConfig {
    conf, _ := Parse()
    return conf.Mail
}

func Yts() YtsConfig {
    conf, _ := Parse()
    return conf.Yts
}

