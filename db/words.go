package db

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"pan_task/stats"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// DB global static DB
var DB map[string][]string

// InitDB inits the DB from the path from config file
func InitDB(dbPath string) error {
	DB = map[string][]string{}
	file, err := os.Open(dbPath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wordsCount := int64(0)
	start := time.Now()

	// Read the file line-by-line and insert its counting array representation string to the db as key,
	// original word will be appended to list of original words as value
	for scanner.Scan() {
		scannedWord := scanner.Text()
		strRep := getStringRepresentationOfCountingArr(getCharCountingArr(scannedWord))
		DB[strRep] = append(DB[strRep], scannedWord)
		wordsCount++
	}

	fmt.Println("[InitDB]: Words scan elapsed", time.Since(start).String())
	fmt.Println("[InitDB]: Total words:", wordsCount)
	stats.SetTotalWords(wordsCount)
	return nil
}

// getCharCountingArr returns a counting array of the word.
//
// For example, for the word 'paloalto' it will return:
//
// a b c d e f g h i j k l m n o p q r s t u v w x y z
//
// 2 0 0 0 0 0 0 0 0 0 0 2 0 0 2 1 0 0 0 1 0 0 0 0 0 0
func getCharCountingArr(word string) *[26]int {
	var arr [26]int
	const AsciiValueOfa = 97
	for _, ch := range word {
		arr[int(ch)-AsciiValueOfa]++
	}

	return &arr
}

// getStringRepresentationOfCountingArr returns a string representation of a counting arr.
// For example, for the word 'paloalto' it will return:
//
// 2:0:0:0:0:0:0:0:0:0:0:2:0:0:2:1:0:0:0:1:0:0:0:0:0:0:
// , where a colon is separating the counter for each character
func getStringRepresentationOfCountingArr(arr *[26]int) string {
	var strRepresentation strings.Builder
	const separator = ":"
	for _, v := range arr {
		strRepresentation.WriteString(strconv.Itoa(v) + separator)
	}

	return strRepresentation.String()
}

// removeElementByValue removes string element from array and returns the array.
//
// If isn't exist, return the array without any change.
func removeElementByValue(arr []string, elem string) []string {
	for i, val := range arr {
		if elem == val {
			if len(arr) > 1 {
				return append(arr[:i], arr[i+1:]...)
			} else {
				// The case where there is only one element on the array, and it is the value itself
				return []string{}
			}
		}
	}

	return arr
}

// GetSimilarWords returns list of permutations of word that exist in the DB
func GetSimilarWords(c *gin.Context) {
	word := c.Query("word")
	if word == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "word parameter value is empty"})
		return
	}

	strRep := getStringRepresentationOfCountingArr(getCharCountingArr(word))
	similarWordsList, ok := DB[strRep]
	// If the key doesn't exist, return empty string array
	if !ok {
		c.IndentedJSON(http.StatusOK, gin.H{"similar": []string{}})
		return
	}

	tempWordsList := make([]string, len(similarWordsList))
	copy(tempWordsList, similarWordsList)
	tempWordsList = removeElementByValue(tempWordsList, word)
	// Return list of words without the word query parameter
	c.JSON(http.StatusOK, gin.H{"similar": tempWordsList})
}
