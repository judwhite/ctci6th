package stringperm

import (
	"testing"
)

// Given a smaller string s and a bigger string b, design an algorithm to find all permutations of the shorter
// string within the longer one. Print the location of each permutation.
//
// Example:
// s: abbc
// b: cbabadcbbabbcbabaabccbabc

func TestBruteForce(t *testing.T) {
	s := "abbc"
	b := "cbabadcbbabbcbabaabccbabc"
	expected := []int{0, 6, 9, 11, 12, 20, 21}
	actual := bruteForce(s, b)
	equal(t, expected, actual)
}

func TestHashMapCompare(t *testing.T) {
	cases := []struct {
		s        string
		b        string
		expected []int
	}{
		{
			s:        "abbc",
			b:        "cbabadcbbabbcbabaabccbabc",
			expected: []int{0, 6, 9, 11, 12, 20, 21},
		},
		{
			s:        "aaaaaaaaaaaaaa",
			b:        "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaa",
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 38, 39, 40, 41},
		},
	}

	for _, c := range cases {
		actual := hashMapCompare(c.s, c.b)
		equal(t, c.expected, actual)
	}
}

func TestSlidingWindow(t *testing.T) {
	cases := []struct {
		s        string
		b        string
		expected []int
	}{
		{
			s:        "abbc",
			b:        "cbabadcbbabbcbabaabccbabc",
			expected: []int{0, 6, 9, 11, 12, 20, 21},
		},
		{
			s:        "aaaaaaaaaaaaaa",
			b:        "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaa",
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 38, 39, 40, 41},
		},
	}

	for _, c := range cases {
		actual := slidingWindowCompare(c.s, c.b)
		equal(t, c.expected, actual)
	}
}

func equal(t *testing.T, expected, actual []int) {
	if len(expected) != len(actual) {
		t.Errorf("want: %v\n\t\t\tgot:  %v", expected, actual)
		return
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("want: %v\n\t\t\tgot:  %v", expected, actual)
			return
		}
	}
}

func BenchmarkBruteForce(b *testing.B) {
	s := "abbc"
	s2 := "cbabadcbbabbcbabaabccbabc"
	for i := 0; i < b.N; i++ {
		bruteForce(s, s2)
	}
}

func BenchmarkHashMapCompare(b *testing.B) {
	s := "abbc"
	s2 := "cbabadcbbabbcbabaabccbabc"
	for i := 0; i < b.N; i++ {
		hashMapCompare(s, s2)
	}
}

func BenchmarkSlidingWindow(b *testing.B) {
	s := "abbc"
	s2 := "cbabadcbbabbcbabaabccbabc"
	for i := 0; i < b.N; i++ {
		slidingWindowCompare(s, s2)
	}
}

func BenchmarkHashMapCompareLong(b *testing.B) {
	s := "aaaaaaaaaaaaaa"
	s2 := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaa"
	for i := 0; i < b.N; i++ {
		hashMapCompare(s, s2)
	}
}

func BenchmarkSlidingWindowLong(b *testing.B) {
	s := "aaaaaaaaaaaaaa"
	s2 := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaa"
	for i := 0; i < b.N; i++ {
		slidingWindowCompare(s, s2)
	}
}
