package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"pan_task/api"
	"pan_task/config"
	"pan_task/db"
	"pan_task/stats"
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
	if err := readConfigFile("config/local.json", config.AppConfig); err != nil {
		panic(err.Error())
	}

	api.InitRouter()

	stats.InitStats()

	if err := db.InitDB(config.AppConfig.DBPath); err != nil {
		panic(err.Error())
	}
}

func main() {
	api.Router.Run(fmt.Sprintf("%s:%s", config.AppConfig.Host, config.AppConfig.Port))
}
