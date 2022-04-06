package stats

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// Singleton like object with small first letter to ensure that the changes on that struct will be only on that package
var programStats *Stats

type Stats struct {
	TotalWords          int64 `json:"totalWords"`
	TotalRequests       int64 `json:"totalRequests"`
	AvgProcessingTimeNs int64 `json:"avgProcessingTimeNs"`
	// RWMutex to allow multiple readers
	lock sync.RWMutex
}

// InitStats inits the stats with zero values
func InitStats() {
	programStats = &Stats{
		TotalWords:          0,
		TotalRequests:       0,
		AvgProcessingTimeNs: 0,
	}
}

// GetStats get statistics on the server
func GetStats(c *gin.Context) {
	programStats.lock.RLock()
	defer programStats.lock.RUnlock()
	c.IndentedJSON(http.StatusOK, programStats)
}

// UpdateStats update statistics after GetSimilarWords endpoint
func UpdateStats(startTime time.Time) {
	programStats.lock.Lock()
	defer programStats.lock.Unlock()
	programStats.TotalRequests += 1
	currentProcessingTimeNs := time.Since(startTime).Nanoseconds()
	programStats.AvgProcessingTimeNs = (programStats.AvgProcessingTimeNs*(programStats.TotalRequests-1) + currentProcessingTimeNs) / programStats.TotalRequests
}

// SetTotalWords set total of words from DB
func SetTotalWords(totalWords int64) {
	programStats.TotalWords = totalWords
}
