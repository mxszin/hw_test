package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type pair struct {
	Key   string
	Value int
}

type pairList []pair

// A function to turn a map into a pairList, then sort and return it.
func sortMapByValue(m map[string]int) pairList {
	p := make(pairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = pair{k, v}
		i++
	}
	sort.Slice(p, func(i, j int) bool {
		if p[i].Value == p[j].Value {
			return p[i].Key < p[j].Key
		}
		return p[i].Value > p[j].Value
	})
	return p
}

var nonWordCharReg = regexp.MustCompile(`[.,!?:;]|(\s-\s)|(\A-)|(-\z)`)

func Top10(text string) []string {
	text = nonWordCharReg.ReplaceAllLiteralString(text, "")
	text = strings.ToLower(text)
	words := strings.Fields(text)

	wordsByCount := make(map[string]int, len(words))
	for _, s := range words {
		wordsByCount[s]++
	}

	sortedWordsAndCounts := sortMapByValue(wordsByCount)

	sortedWords := make([]string, 0, 10)
	for i := 0; i < len(sortedWordsAndCounts); i++ {
		p := sortedWordsAndCounts[i]
		sortedWords = append(sortedWords, p.Key)

		if i >= 9 {
			break
		}
	}

	return sortedWords
}
