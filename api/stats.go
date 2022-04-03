package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var ProgramStats Stats

type Stats struct {
	TotalWords          int `json:"totalWords"`
	TotalRequests       int `json:"totalRequests"`
	AvgProcessingTimeNs int `json:"avgProcessingTimeNs"`
}

func GetStats(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ProgramStats)
}
