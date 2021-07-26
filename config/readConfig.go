package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Config struct {
	Gotify struct {
		URL      string `yaml:"url"`
		ApiToken string `yaml:"apiToken"`
	}
	Matrix struct {
		HomeServerURL string `yaml:"homeserverURL"`
		Username      string `yaml:"username"`
		Password      string `yaml:"password"`
		Token         string `yaml:"token"`
		RoomID        string `yaml:"roomID"`
	}
}

func readConf() *Config {
	buf, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Fatal("Could not load config. ", err)
	}

	c := &Config{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		log.Fatal("Could not load config.", err)
	}

	return c
}

var Configuration = readConf()
