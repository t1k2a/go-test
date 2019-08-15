package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Setting struct {
	Title       string `yaml:"title"`
	BaseUrl     string `yaml:"base_url"`
	MailAddress string `yaml:"mail_address"`
	Password    string `yaml:"password"`
	OutputPath  string `yaml:"output_path"`
	NotifyUlr   string `yaml:"notify_url"`
}

func LoadFile(environment string) *Setting {
	buf, err := ioutil.ReadFile("./settings/" + environment + ".yaml")
	if err != nil {
		panic(err)
	}

	s := new(Setting)
	err = yaml.Unmarshal(buf, &s)
	if err != nil {
		panic(err)
	}

	return s
}
