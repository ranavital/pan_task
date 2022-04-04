package stats

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var ProgramStats *Stats

type Stats struct {
	TotalWords          int64        `json:"totalWords"`
	TotalRequests       int64        `json:"totalRequests"`
	AvgProcessingTimeNs int64        `json:"avgProcessingTimeNs"`
	lock                sync.RWMutex `json:"-"`
}

func InitStats() {
	ProgramStats = &Stats{
		TotalWords:          0,
		TotalRequests:       0,
		AvgProcessingTimeNs: 0,
	}
}

func GetStats(c *gin.Context) {
	ProgramStats.lock.RLock()
	defer ProgramStats.lock.RUnlock()
	c.IndentedJSON(http.StatusOK, ProgramStats)
}

func UpdateStats(currentAvgProcessingTimeNs int64) {
	ProgramStats.lock.Lock()
	defer ProgramStats.lock.Unlock()
	ProgramStats.AvgProcessingTimeNs = (ProgramStats.AvgProcessingTimeNs*ProgramStats.TotalRequests + currentAvgProcessingTimeNs) / (ProgramStats.TotalRequests + 1)
	ProgramStats.TotalRequests += 1
}

func SetTotalWords(totalWords int64) {
	ProgramStats.TotalWords = totalWords
}
