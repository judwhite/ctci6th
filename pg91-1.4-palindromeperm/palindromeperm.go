package palindromeperm

import "unicode"

func isPalindromePermutationASCII(s string) bool {
	m := make([]byte, 256)
	diffs := 0
	lenLessSpaces := 0
	for i := 0; i < len(s); i++ {
		c := unicode.ToLower(rune(s[i]))
		if c == ' ' {
			continue
		}
		lenLessSpaces++

		v := m[c]
		if v == 0 {
			diffs++
			m[c] = 1
		} else {
			diffs--
			m[c] = 0
		}
	}

	return diffs == (lenLessSpaces & 1)
}

func isPalindromePermutationUnicode(s string) bool {
	m := make(map[rune]byte)
	diffs := 0
	lenLessSpaces := 0
	for _, c := range s {
		if c == ' ' {
			continue
		}
		c = unicode.ToLower(c)
		lenLessSpaces++

		v := m[c]
		if v == 0 {
			diffs++
			m[c] = 1
		} else {
			diffs--
			m[c] = 0
		}
	}

	return diffs == (lenLessSpaces & 1)
}
