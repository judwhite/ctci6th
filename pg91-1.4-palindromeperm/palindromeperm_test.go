package palindromeperm

import (
	"testing"
)

var cases = []struct {
	input    string
	isASCII  bool
	expected bool
}{
	{"Tact Coa", true, true},
	{"Tact Coat", true, false},
	{"Tact Coaz", true, false},
	{"ABcba", true, true},
	{"ABcCba", true, true},
	{"ABba", true, true},
	{"AB界ba", false, true},
	{"AB界界ba", false, true},
	{"AB界aa", false, false},
	{"AB界界baz", false, true},
	{"a  界 界  A", false, true},
	{"a  界 界  a a", false, true},
	{"a  界 界  a b", false, true},
	{"a  界 界  b", false, false},
}

func TestIsPalindromePermutationASCII(t *testing.T) {
	for _, c := range cases {
		if !c.isASCII {
			continue
		}
		actual := isPalindromePermutationASCII(c.input)
		if actual != c.expected {
			t.Errorf("input: %q want: %v got: %v", c.input, c.expected, actual)
		}
	}
}

func TestIsPalindromePermutationUnicode(t *testing.T) {
	for _, c := range cases {
		actual := isPalindromePermutationUnicode(c.input)
		if actual != c.expected {
			t.Errorf("input: %q want: %v got: %v", c.input, c.expected, actual)
		}
	}
}

func BenchmarkIsPalindromePermutationASCII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindromePermutationASCII("Tact Coa")
	}
}

func BenchmarkIsPalindromePermutationUnicode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindromePermutationUnicode("Tact Coa")
	}
}
