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
		PollTime string `yaml:"pollTime"`
	}
	Matrix struct {
		HomeServerURL string `yaml:"homeserverURL"`
		Username      string `yaml:"username"`
		Token         string `yaml:"token"`
		RoomID        string `yaml:"roomID"`
		Encrypted     bool   `yaml:"encrypted"`
	}
	Debug bool `yaml:"debug"`
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

func checkValues(config *Config) {

	if config.Gotify.URL == "" {
		log.Fatal("No gotify url specified.")
	}

	if config.Gotify.ApiToken == "" {
		log.Fatal("No gotify api token specified.")
	}

	if config.Gotify.PollTime == "" {
		log.Fatal("No gotify polltime specified.")
	}

	if config.Matrix.HomeServerURL == "" {
		log.Fatal("No matrix homeserver specified.")
	}

	if config.Matrix.Username == "" {
		log.Fatal("No matrix username specified.")
	}

	if config.Matrix.Token == "" {
		log.Fatal("No matrix auth token specified.")
	}

	if config.Matrix.RoomID == "" {
		log.Fatal("No matrix room id specified.")
	}

	if !config.Matrix.Encrypted {
		log.Fatal("No encryption specified. Please use true or false")
	}

}
