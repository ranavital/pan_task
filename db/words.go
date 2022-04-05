package db

import (
	"bufio"
	"net/http"
	"os"
	"pan_task/stats"
	"sort"

	"github.com/gin-gonic/gin"
)

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

var DB map[string][]string

func InitDB(dbPath string) error {
	DB = map[string][]string{}
	file, err := os.Open(dbPath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wordsCount := int64(0)
	for scanner.Scan() {
		scannedWord := scanner.Text()
		sortedWord := sortString(scannedWord)
		DB[sortedWord] = append(DB[sortedWord], scannedWord)
		wordsCount++
	}

	stats.SetTotalWords(wordsCount)
	return nil
}

//FIX BUG IN FUNCTION
func removeElementByValue(arr []string, elem string) []string {
	for i, val := range arr {
		if elem == val {
			if len(arr) > 1 {
				// Secure last element
				return append(arr[:i], arr[i+1:]...)
			} else {
				// The case where there is only one element on the array and it is the value itself
				return []string{}
			}
		}
	}

	return arr
}

func GetSimilarWords(c *gin.Context) {
	word := c.Query("word")
	similarWordsList, ok := DB[sortString(word)]
	if !ok {
		c.IndentedJSON(http.StatusOK, gin.H{"similar": []string{}})
		return
	}
	tempWordsList := make([]string, len(similarWordsList))
	copy(tempWordsList, similarWordsList)
	tempWordsList = removeElementByValue(tempWordsList, word)
	c.JSON(http.StatusOK, gin.H{"similar": tempWordsList})
}
