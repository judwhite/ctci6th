package stringrot

import (
	"testing"
)

var cases = []struct {
	s1       string
	s2       string
	expected bool
}{
	{
		s1:       "erbottlewat",
		s2:       "waterbottle",
		expected: true,
	},
	{
		s1:       "waterbottle",
		s2:       "erbottlewat",
		expected: true,
	},
	{
		s1:       "aaaaaaaaaab",
		s2:       "baaaaaaaaaa",
		expected: true,
	},
	{
		s1:       "aaaaaaaaaab",
		s2:       "aaaaaaaaaba",
		expected: true,
	},
	{
		s1:       "aaaaaaaaaba",
		s2:       "aaaaaaaaaab",
		expected: true,
	},
	{
		s1:       "aaaaaaaaaab",
		s2:       "aaaaaaaaabb",
		expected: false,
	},
	{
		s1:       "waterbottlewaterbottle",
		s2:       "erbottlewat",
		expected: false,
	},
	{
		// cause the stdlib to use Rabin-Karp search by exceeding const shortStringLen = 31
		s1:       "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		s2:       "Zabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXY",
		expected: true,
	},
}

func TestIsRotationBruteForce(t *testing.T) {
	for _, c := range cases {
		actual := isRotationBruteForce(c.s1, c.s2)
		if actual != c.expected {
			t.Errorf("s1: %q s2: %q want: %v got: %v", c.s1, c.s2, c.expected, actual)
		}
	}
}

func TestIsRotationSubstring(t *testing.T) {
	for _, c := range cases {
		actual := isRotationSubstring(c.s1, c.s2)
		if actual != c.expected {
			t.Errorf("s1: %q s2: %q want: %v got: %v", c.s1, c.s2, c.expected, actual)
		}
	}
}

func BenchmarkIsRotationBruteForce(b *testing.B) {
	for _, c := range cases {
		b.Run(trimName(c.s1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isRotationBruteForce(c.s1, c.s2)
			}
		})
	}
}

func BenchmarkIsRotationSubstring(b *testing.B) {
	for _, c := range cases {
		b.Run(trimName(c.s1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isRotationSubstring(c.s1, c.s2)
			}
		})
	}
}

func trimName(s string) string {
	if len(s) <= 25 {
		return s
	}
	return s[:22] + "..."
}
