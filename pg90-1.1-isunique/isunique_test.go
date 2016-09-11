package isunique

import (
	"testing"
)

var cases = []struct {
	input    string
	isASCII  bool
	expected bool
}{
	{"", true, true},
	{"a", true, true},
	{"abc", true, true},
	{"baeac", true, false},
	{"zazb", true, false},
	{"bzaz", true, false},
	{"zbza", true, false},
	{"zzba", true, false},
	{"bazz", true, false},
	{"Helo, 世界", false, true},
	{"界世界", false, false},
}

func TestIsUniqueSort(t *testing.T) {
	for _, c := range cases {
		actual := isUniqueSort(c.input)
		if actual != c.expected {
			t.Errorf("%q want: %v got: %v", c.input, c.expected, actual)
		}
	}
}

func TestIsUniqueArray(t *testing.T) {
	for _, c := range cases {
		if !c.isASCII {
			// isUniqueArray only supports ASCII characters
			continue
		}
		actual := isUniqueArray(c.input)
		if actual != c.expected {
			t.Errorf("%q want: %v got: %v", c.input, c.expected, actual)
		}
	}
}

func TestIsUniqueMap(t *testing.T) {
	for _, c := range cases {
		actual := isUniqueMap(c.input)
		if actual != c.expected {
			t.Errorf("%q want: %v got: %v", c.input, c.expected, actual)
		}
	}
}

func BenchmarkIsUniqueSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isUniqueSort("zyxwvutsrqponmlkjihgfedcba")
	}
}

func BenchmarkIsUniqueArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isUniqueArray("zyxwvutsrqponmlkjihgfedcba")
	}
}

func BenchmarkIsUniqueMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isUniqueMap("zyxwvutsrqponmlkjihgfedcba")
	}
}
