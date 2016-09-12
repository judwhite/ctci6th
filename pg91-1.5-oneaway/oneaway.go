package oneaway

import "unicode/utf8"

func isOneAway(s1, s2 string) bool {
	lenDiff := utf8.RuneCountInString(s1) - utf8.RuneCountInString(s2)
	if lenDiff < 0 {
		s1, s2 = s2, s1 // make the longer string s1, simplifies the loop
	}
	if lenDiff > 1 || lenDiff < -1 {
		return false
	}

	i, j, diff := 0, 0, 0
	for j < len(s2) {
		c1, n1 := utf8.DecodeRuneInString(s1[i:])
		i += n1
		c2, n2 := utf8.DecodeRuneInString(s2[j:])
		j += n2

		// match
		if c1 == c2 {
			continue
		}

		diff++
		if diff > 1 {
			return false
		}

		if lenDiff != 0 {
			// read next rune from s1
			c1, n1 := utf8.DecodeRuneInString(s1[i:])
			if c1 != c2 {
				return false
			}
			i += n1
		}
	}

	return true
}
