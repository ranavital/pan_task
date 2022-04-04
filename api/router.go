package api

import (
	"pan_task/db"
	"pan_task/stats"
	"time"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitRouter() {
	Router = gin.Default()
	apiV1 := Router.Group("/api/v1")
	{
		apiV1.GET("/stats", stats.GetStats)
		apiV1.GET("/similar", calcProcessingTime, db.GetSimilarWords)
	}
}

func calcProcessingTime(c *gin.Context) {
	now := time.Now()
	c.Next()
	stats.UpdateStats(time.Since(now).Nanoseconds())
}
