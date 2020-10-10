package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"regexp"
	"sort"
	"strings"
)

type wordCount struct {
	word  string
	count int
}

func Top10(str string) []string {
	a := regexp.MustCompile(`(\s|[,:!."';])`)
	m := make(map[string]int)
	for _, s := range a.Split(str, -1) {
		if s == "" {
			continue
		}

		s = strings.ToLower(s)

		if num, ok := m[s]; ok {
			m[s] = num + 1
		} else {
			m[s] = 1
		}
	}

	var prepList []wordCount
	for w, c := range m {
		prepList = append(prepList, wordCount{w, c})
	}

	sort.Slice(prepList, func(i int, j int) bool {
		return prepList[i].count > prepList[j].count
	})

	var result []string
	for i, wc := range prepList {
		if wc.word == "-" {
			continue
		}
		result = append(result, wc.word)
		if i == 9 {
			break
		}
	}
	return result
}
