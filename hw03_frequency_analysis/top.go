package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var Replacer = regexp.MustCompile(`[\s.,!?"]+`)

type WordCounter struct {
	word  string
	count int
}

func Top10(str string) []string {
	position := 10
	str = Replacer.ReplaceAllString(str, " ")
	wordsRaw := strings.Split(str, " ")
	wordsCount := map[string]int{}
	for _, word := range wordsRaw {
		if word == "-" {
			continue
		}
		wordsCount[strings.ToLower(word)]++
	}
	wordCounters := make([]WordCounter, 0)
	for word, count := range wordsCount {
		if word == "" || word == "-" {
			continue
		}
		wordCounters = append(wordCounters, WordCounter{word, count})
	}

	sort.Slice(wordCounters, func(i, j int) bool {
		if wordCounters[i].count == wordCounters[j].count {
			return wordCounters[i].word < wordCounters[j].word
		}
		return wordCounters[i].count > wordCounters[j].count
	})
	top10 := make([]string, 0)
	if length := len(wordCounters); length < position {
		position = length
	}
	for _, word := range wordCounters[:position] {
		top10 = append(top10, word.word)
	}

	return top10
}
