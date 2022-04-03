package api

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine{
	router := gin.Default()
	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/stats", GetStats)
	}

	return router
}
