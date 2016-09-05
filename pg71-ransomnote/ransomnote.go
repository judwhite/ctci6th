package ransomnote

import (
	"strings"
	"unicode"
)

func isRansomNote(note, corpus string) bool {
	m1 := getWordMap(corpus)
	m2 := getWordMap(note)

	if len(m2) > len(m1) {
		return false
	}

	for k, v := range m2 {
		if v2, ok := m1[k]; !ok {
			return false
		} else {
			if v2 < v {
				return false
			}
		}
	}
	return true
}

func getWordMap(corpus string) map[string]int {
	corpus = strings.Replace(corpus, "\n", " ", -1)
	corpus = strings.Replace(corpus, "\r", " ", -1)
	words := strings.Split(corpus, " ")

	m := make(map[string]int)
	for i := 0; i < len(words); i++ {
		word := strings.ToLower(words[i])
		word = strings.TrimFunc(word, unicode.IsPunct)
		if len(word) == 0 {
			continue
		}
		words[i] = word
		m[word]++
	}
	return m
}
