package oneaway

import (
	"testing"
)

var cases = []struct {
	a        string
	b        string
	expected bool
}{
	{"pale", "ple", true},
	{"pales", "pale", true},
	{"pale", "bale", true},
	{"pale", "bake", false},
	{"pale", "pakes", false},
	{"pale", "pale", true},
	{"ple", "pale", true},
	{"pale", "pales", true},
	{"bale", "pale", true},
	{"bale", "pake", false},
	{"pakes", "pale", false},
	{"pale", "palee", true},
	{"pale", "paleee", false},
	{"palee", "pale", true},
	{"paleee", "pale", false},
}

func TestIsOneAway(t *testing.T) {
	for _, c := range cases {
		actual := isOneAway(c.a, c.b)
		if actual != c.expected {
			t.Errorf("a: %q b: %q want: %v got: %v", c.a, c.b, c.expected, actual)
		}
	}
}

func BenchmarkIsOneAway(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isOneAway("pale", "pales")
	}
}
