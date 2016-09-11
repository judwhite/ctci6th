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
	var m [256]bool
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

// isUniqueBitVector returns true if all characters in a string are unique, using a bit vector
// This solution I didn't initially think of. I checked the solutions on pg192 and saw "bit vector" and went back
// to the code to implement it. It'll only work with ASCII characters which usually isn't a valid assumption, but
// solving the general problem of uniqueness in an array given a a bounded range of values is interesting.
// Update: The benchmark is about twice as slow as the isUniqueArray function, probably due to the additional
// math. The memory usage is 4 bytes so this technique certainly has application. The solution in the book assumed
// lowercase a-z to use a single int, which would reduce the math.
func isUniqueBitVector(s string) bool {
	var m [4]uint64
	for i := 0; i < len(s); i++ {
		c := uint64(s[i])
		bit := uint64(1 << (c & 0x3F))
		index := c >> 6
		if m[index]&bit > 0 {
			return false
		}
		m[index] |= bit
	}
	return true
}

// isUniqueBitVector2 returns true if all characters in a string are unique, using a bit vector
// it's limited to lowercase a-z. in a cruel twist of fate, the 256-byte array version is still faster and
// covers more of the ASCII set, though this uses less memory. always run your benchmarks.
func isUniqueBitVector2(s string) bool {
	var m uint
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		bit := uint(1 << c)
		if m&bit > 0 {
			return false
		}
		m |= bit
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
