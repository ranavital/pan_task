package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"pan_task/api"
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
	if err := readConfigFile("config/local.json", config.AppConfig); err != nil {
		panic(err.Error())
	}

	api.ProgramStats = api.Stats{
		TotalWords:          0,
		TotalRequests:       0,
		AvgProcessingTimeNs: 0,
	}
}

func main() {
	router := api.InitRouter()
	router.Run(fmt.Sprintf("%s:%s", config.AppConfig.Host, config.AppConfig.Port))
}
