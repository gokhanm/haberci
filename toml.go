package main

import (
	"github.com/BurntSushi/toml"
)

var configPath string

// TomlConfig global toml config struct
type TomlConfig struct {
	Title string

	Mail struct {
		Server   string
		Port     int
		Username string
		Password string
		From     string
	}

	Yts struct {
		Enabled       bool
		Subject       string
		ToRecipients  []string
		BccRecipients []string
		PageLimit     string
	}
}

// Load given "c" flag conf file
func Load(conf string) {
	configPath = conf
}

// Parse is parse toml config file
func Parse() (*TomlConfig, error) {
	var conf *TomlConfig

	if _, err := toml.DecodeFile(configPath, &conf); err != nil {
		return nil, err
	}

	return conf, nil
}
