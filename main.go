package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"pan_task/config"
)

func readConfigFile(path string, conf *config.Config) error {
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

func init() {
	err := readConfigFile("config/local.json", config.AppConfig)
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	fmt.Printf("%s:%s\n", config.AppConfig.Host, config.AppConfig.Port)
}
