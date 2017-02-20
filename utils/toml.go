package toml

import (
    "errors"

    "github.com/BurntSushi/toml"
)


type MailConfig struct {
    Server string   `toml:"server"`
    Port int        `toml:"port"`
    Username string `toml:"username"`
    Password string `toml:"password"`
    From string     `toml:"from"`
}

type YtsConfig struct{
    PageLimit string    `toml:"pageLimit"`
    Recipients []string `toml:"recipients"`
    Subject string      `toml:"subject"`
}

type tConf struct {
    Title string
    Mail MailConfig
    Yts YtsConfig
}

func Parse() (*tConf, error) {
    tomlData := "/etc/haberci.toml"

    var conf *tConf

    if _, err := toml.DecodeFile(tomlData, &conf); err != nil {
        panic(errors.New("haberci.toml conf file not found under /etc folder"))
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

