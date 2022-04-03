package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"pan_task/api"
	"pan_task/config"
	"pan_task/db"
	"path"
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

	if err := db.InitDB(path.Join("..", config.AppConfig.DBPath)); err != nil {
		panic(err.Error())
	}

	api.ProgramStats = api.Stats{
		TotalWords:          0,
		TotalRequests:       0,
		AvgProcessingTimeNs: 0,
	}
}

func main() {
	api.Router.Run(fmt.Sprintf("%s:%s", config.AppConfig.Host, config.AppConfig.Port))
}
