package isunique

import (
	"sort"
)

// Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you cannot
// use additional data structures?

// isUniqueSort returns true if all characters in a string are unique, using a sorting algorithm
func isUniqueSort(s string) bool {
	if len(s) <= 1 {
		return true
	}
	r := runeSlice(s)
	sort.Sort(r)
	for i := 1; i < len(r); i++ {
		if r[i] == r[i-1] {
			return false
		}
	}
	return true
}

// isUniqueArray returns true if all characters in a string are unique, using an array (limited to ASCII characters)
func isUniqueArray(s string) bool {
	m := make([]bool, 256)
	for i := 0; i < len(s); i++ {
		v := s[i]
		if m[v] {
			return false
		}
		m[v] = true
	}
	return true
}

// isUniqueMap returns true if all characters in a string are unique, using a map
func isUniqueMap(s string) bool {
	m := make(map[rune]struct{}, 256)
	for _, r := range s {
		_, ok := m[r]
		if ok {
			return false
		}
		m[r] = struct{}{}
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
