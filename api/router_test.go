package api_test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"pan_task/api"
	"pan_task/config"
	"pan_task/db"
	"pan_task/stats"
	"sort"
	"testing"
	"time"
)

// String sort functionality
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func init() {
	if err := config.ReadConfigFile("../config/test.json"); err != nil {
		panic("[init]: failed to read config file: " + err.Error())
	}

	api.InitRouter(gin.TestMode)
	stats.InitStats()
	if err := db.InitDB(config.AppConfig.DBPath); err != nil {
		panic("[init]: failed to init db: " + err.Error())
	}

	rand.Seed(time.Now().UnixNano())
}

type similarResponse struct {
	Similar []string `json:"similar"`
}

var similarWordsList = []string{"apple", "aba", "aab", "baa", "stressed", "pan", "nap", "hrta", "abbe", "rabb", "abear"}

// For the total requests stats test
var totalRequestsSoFar = 0

const iterations = 100

func Test_server(t *testing.T) {
	var word string
	var req *http.Request
	var w *httptest.ResponseRecorder
	var similarRes similarResponse
	for i := 0; i < iterations; i++ {
		word = randStringRunes()
		w = httptest.NewRecorder()
		req = httptest.NewRequest("GET", "/api/v1/similar?word="+word, nil)
		api.Router.ServeHTTP(w, req)
		assert.Equal(t, w.Code, http.StatusOK)
		err := json.Unmarshal(w.Body.Bytes(), &similarRes)
		assert.NoError(t, err)
		sortedWord := sortString(word)
		for _, perm := range similarRes.Similar {
			assert.Equal(t, sortedWord, sortString(perm))
		}
		totalRequestsSoFar++
	}

	testStats(t, totalRequestsSoFar)
}

func Test_server_similar_words_list(t *testing.T) {
	var req *http.Request
	var w *httptest.ResponseRecorder
	var similarRes similarResponse
	for _, word := range similarWordsList {
		w = httptest.NewRecorder()
		req = httptest.NewRequest("GET", "/api/v1/similar?word="+word, nil)
		api.Router.ServeHTTP(w, req)
		assert.Equal(t, w.Code, http.StatusOK)
		err := json.Unmarshal(w.Body.Bytes(), &similarRes)
		assert.NoError(t, err)
		sortedWord := sortString(word)
		for _, perm := range similarRes.Similar {
			assert.Equal(t, sortedWord, sortString(perm))
		}
		totalRequestsSoFar++
	}

	testStats(t, totalRequestsSoFar)
}

func testStats(t *testing.T, totalRequests int) {
	w := httptest.NewRecorder()
	statsReq := httptest.NewRequest("GET", "/api/v1/stats", nil)
	api.Router.ServeHTTP(w, statsReq)
	var statsRes stats.Stats
	err := json.Unmarshal(w.Body.Bytes(), &statsRes)
	assert.NoError(t, err)
	assert.Equal(t, int(statsRes.TotalRequests), totalRequests)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func randStringRunes() string {
	strLength := rand.Intn(4) + 3 // Maximum word length range - [3,6]
	b := make([]rune, strLength)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Test_ReadConfigFile_false_path(t *testing.T) {
	err := config.ReadConfigFile("false_path")
	assert.Error(t, err)
}

func Test_InitDB_db_false_path(t *testing.T) {
	err := db.InitDB("false_path")
	assert.Error(t, err)
}

func Benchmark_server(b *testing.B) {
	var word string
	var req *http.Request
	var w *httptest.ResponseRecorder
	var similarRes similarResponse
	for i := 0; i < b.N; i++ {
		word = randStringRunes()
		w = httptest.NewRecorder()
		req = httptest.NewRequest("GET", "/api/v1/similar?word="+word, nil)
		api.Router.ServeHTTP(w, req)
		assert.Equal(b, w.Code, http.StatusOK)
		err := json.Unmarshal(w.Body.Bytes(), &similarRes)
		assert.NoError(b, err)
		sortedWord := sortString(word)
		for _, perm := range similarRes.Similar {
			assert.Equal(b, sortedWord, sortString(perm))
		}
		totalRequestsSoFar++
	}

	benchmarkTestStats(b, totalRequestsSoFar)
}

func benchmarkTestStats(b *testing.B, totalRequests int) {
	w := httptest.NewRecorder()
	statsReq := httptest.NewRequest("GET", "/api/v1/stats", nil)
	api.Router.ServeHTTP(w, statsReq)
	var statsRes stats.Stats
	err := json.Unmarshal(w.Body.Bytes(), &statsRes)
	assert.NoError(b, err)
	assert.Equal(b, int(statsRes.TotalRequests), totalRequests)
}
