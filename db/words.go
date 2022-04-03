package db

import (
	"bufio"
	"log"
	"os"
	"sort"
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
	for scanner.Scan() {
		scannedWord := scanner.Text()
		sortedWord := sortString(scannedWord)
		DB[sortedWord] = append(DB[sortedWord], scannedWord)
	}
}
