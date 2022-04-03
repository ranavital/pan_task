package api

import "github.com/gin-gonic/gin"

var Router *gin.Engine

func InitRouter() {
	Router = gin.Default()
	apiV1 := Router.Group("/api/v1")
	{
		apiV1.GET("/stats", GetStats)
		// apiV1.GET("/similar", GetSimilarWords)
	}
}
