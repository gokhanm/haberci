package toml

import "github.com/BurntSushi/toml"


type MailConfig struct {
    Server string   `toml:"server"`
    Port int        `toml:"port"`
    Username string `toml:"username"`
    Password string `toml:"password"`
    From string     `toml:"from"`
    To []string     `toml:"to"`
    Subject string  `toml:"subject"`
}

type YtsConfig struct{
    PageLimit string `toml:"pageLimit"`
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
        return nil, err
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

