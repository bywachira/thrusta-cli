package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config defines the cli config
type Config struct {
	SSL  bool   `json:"ssl"`
	URL  string `json:"url"`
	Path string `json:"path"`
}

// ReadConfig read the config.json file for url and SSL enable
func ReadConfig() Config {
	file, _ := ioutil.ReadFile("config.json")

	data := Config{}

	_ = json.Unmarshal([]byte(file), &data)

	return data
}

// ProtocolChecker handles protocol checkers
func ProtocolChecker() string {
	config := ReadConfig()

	if config.SSL {
		return "wss"
	} else {
		return "ws"
	}
}
