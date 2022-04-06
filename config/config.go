package config

import (
	"encoding/json"
	"io/ioutil"
)

var AppConfig = &Config{}

type Config struct {
	Host   string `json:"host"`
	Port   string `json:"port"`
	DBPath string `json:"db_path"`
}

// ReadConfigFile reads a config json structure file into a conf parameter
func ReadConfigFile(path string, conf *Config) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, conf)
	if err != nil {
		return err
	}

	return nil
}
