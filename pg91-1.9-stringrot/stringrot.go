package stringrot

import (
	"strings"
)

func isRotationSubstring(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	return strings.Index(s1+s1, s2) != -1
}

func isRotationBruteForce(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		found := true
		for j := 0; j < len(s2); j++ {
			k := (i + j) % len(s1)
			if s2[j] != s1[k] {
				found = false
				break
			}
		}
		if found {
			return true
		}
	}
	return false
}
