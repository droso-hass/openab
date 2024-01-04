package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var ConfigData Config

type Config struct {
	Tokens []string `yaml:"tokens,flow"`
	XMPP   struct {
		Port int `yaml:"port"`
	} `yaml:"xmpp,flow"`
}

func Parse(path string) error {
	ConfigData = Config{}
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(file, &ConfigData)
}
