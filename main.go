package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pan_task/api"
	"pan_task/config"
	"pan_task/db"
	"pan_task/stats"
	"runtime"
)

func init() {
	// Init the max processes to the maximum processes usable
	fmt.Println("[init]: setting max processes to:", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("[init]: reading config file")
	if err := config.ReadConfigFile("config/local.json", config.AppConfig); err != nil {
		panic("[init]: failed to read config file: " + err.Error())
	}
	fmt.Println("[init]: successfully read config file")

	fmt.Println("[init]: initializing router")
	api.InitRouter(gin.ReleaseMode)
	fmt.Println("[init]: successfully initialized router")

	fmt.Println("[init]: initializing stats")
	stats.InitStats()
	fmt.Println("[init]: successfully initialized stats")

	fmt.Println("[init]: initializing DB")
	if err := db.InitDB(config.AppConfig.DBPath); err != nil {
		panic("[init]: failed to init db: " + err.Error())
	}
	fmt.Println("[init]: successfully initialized db")
}

func main() {
	if err := api.Router.Run(fmt.Sprintf("%s:%s", config.AppConfig.Host, config.AppConfig.Port)); err != nil {
		panic("[main]: failed to run server: " + err.Error())
	}
}
