package api

import (
	"pan_task/db"
	"pan_task/stats"
	"time"

	"github.com/gin-gonic/gin"
)

// Router global router
var Router *gin.Engine

// InitRouter inits Gin router with api v1 group of 2 endpoints on Release mode
func InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	Router = gin.Default()
	apiV1 := Router.Group("/api/v1")
	{
		apiV1.GET("/stats", stats.GetStats)
		apiV1.GET("/similar", calcProcessingTime, db.GetSimilarWords)
	}
}

// calcProcessingTime middleware for calculating GetSimilarWords request latency and update statistics for that request
func calcProcessingTime(c *gin.Context) {
	now := time.Now()
	c.Next()
	stats.UpdateStats(now)
}
