package checkperm

import (
	"sort"
)

// isPermMap returns true if a is a permutation of b
// returns true if both strings are empty
// supports full unicode
func isPermMap(a, b string) bool {
	ra, rb := []rune(a), []rune(b)
	if len(ra) != len(rb) {
		return false
	}

	m := make(map[rune]int)
	diff := 0
	for i := 0; i < len(ra); i++ {
		// same rune at same pos, skip the math
		if ra[i] == rb[i] {
			continue
		}

		v1 := m[ra[i]] + 1
		v2 := m[rb[i]] - 1
		m[ra[i]] = v1
		m[rb[i]] = v2

		if v1 <= 0 {
			diff--
		} else {
			diff++
		}

		if v2 >= 0 {
			diff--
		} else {
			diff++
		}
	}

	return diff == 0
}

// isPermArray returns true if a is a permutation of b
// returns true if both strings are empty
// supports 1 byte characters
func isPermArray(a, b string) bool {
	ra, rb := []rune(a), []rune(b)
	if len(ra) != len(rb) {
		return false
	}

	var m [2 << 8]int32
	diff := 0
	for i := 0; i < len(ra); i++ {
		// same rune at same pos, skip the math
		if ra[i] == rb[i] {
			continue
		}

		v1 := m[ra[i]] + 1
		v2 := m[rb[i]] - 1
		m[ra[i]] = v1
		m[rb[i]] = v2

		if v1 <= 0 {
			diff--
		} else {
			diff++
		}

		if v2 >= 0 {
			diff--
		} else {
			diff++
		}
	}

	return diff == 0
}

// isPermSort returns true if a is a permutation of b
// returns true if both strings are empty
// supports full unicode
func isPermSort(a, b string) bool {
	ra, rb := runeSlice(a), runeSlice(b)
	if len(ra) != len(rb) {
		return false
	}

	sort.Sort(ra)
	sort.Sort(rb)

	for i := 0; i < len(ra); i++ {
		if ra[i] != rb[i] {
			return false
		}
	}

	return true
}

type runeSlice []rune

func (r runeSlice) Len() int {
	return len(r)
}

func (r runeSlice) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r runeSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
