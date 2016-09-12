package stringcompr

import "strconv"

// String Compression: Implement a method to perform basic string compression using the counts of repeated
// characters. For example, the string aabcccccaaa would become a2b1c5a3. If the "compressed" string would not
// become smaller than the original string, your method should return the original string. You can assume the
// string has only uppercase and lowercase letters (a-z).

func compress(s string) string {
	if len(s) <= 2 {
		return s
	}

	result := make([]byte, len(s)-1) // the resulting string must be smaller than the original
	prev := s[0]
	prevCount := 1
	var c byte
	j := 0

	addNumber := func() bool {
		result[j] = prev
		j++
		if prevCount > 9 {
			prevCountString := strconv.Itoa(prevCount)
			if j+len(prevCountString) > len(result) {
				return false
			}
			copy(result[j:], prevCountString)
			j += len(prevCountString)
		} else {
			if j+1 > len(result) {
				return false
			}
			result[j] = byte(prevCount + 48)
			j++
		}
		prev = c
		prevCount = 1
		return true
	}

	for i := 1; i < len(s); i++ {
		c = s[i]
		if c == prev {
			prevCount++
		} else {
			if !addNumber() {
				return s
			}
		}
	}
	if !addNumber() {
		return s
	}

	return string(result[:j])
}

func compressCountFirst(s string) string {
	if len(s) <= 2 {
		return s
	}

	prev := s[0]
	repeated := 1
	compressedSize := 0
	for i := 1; i < len(s); i++ {
		c := s[i]
		if c == prev {
			repeated++
		} else {
			var repeatedStringCount int
			if repeated > 9 {
				repeatedStringCount = len(strconv.Itoa(repeated))
			} else {
				repeatedStringCount = 1
			}
			compressedSize += 1 + repeatedStringCount
			if compressedSize+2 >= len(s) {
				return s
			}
			prev = c
			repeated = 1
		}
	}

	var repeatedStringCount int
	if repeated > 9 {
		repeatedStringCount = len(strconv.Itoa(repeated))
	} else {
		repeatedStringCount = 1
	}
	compressedSize += 1 + repeatedStringCount
	if compressedSize >= len(s) {
		return s
	}

	b := make([]byte, compressedSize)
	prev = s[0]
	repeated = 1
	j := 0
	for i := 1; i < len(s); i++ {
		c := s[i]
		if c == prev {
			repeated++
		} else {
			b[j] = prev
			if repeated > 9 {
				prevCountString := strconv.Itoa(repeated)
				copy(b[j+1:], prevCountString)
				j += len(prevCountString) + 1
			} else {
				b[j+1] = byte(repeated + 48)
				j += 2
			}
			prev = c
			repeated = 1
		}
	}
	b[j] = prev
	if repeated > 9 {
		prevCountString := strconv.Itoa(repeated)
		copy(b[j+1:], prevCountString)
	} else {
		b[j+1] = byte(repeated + 48)
	}

	return string(b)
}
