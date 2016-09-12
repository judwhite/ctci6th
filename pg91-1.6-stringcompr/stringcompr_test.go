package stringcompr

import "testing"

var cases = []struct {
	input    string
	expected string
}{
	{"", ""},
	{"a", "a"},
	{"aa", "aa"},
	{"aaa", "a3"},
	{"aabcccccaaa", "a2b1c5a3"},
	{"abcd", "abcd"},
	{"aabcccccaaabcdefgh", "aabcccccaaabcdefgh"},
	{"bcdefghaabcccccaaa", "bcdefghaabcccccaaa"},
	{"aabaaaaaaaaaa", "a2b1a10"},
	{"aabcaaaaaaaaaa", "a2b1c1a10"},
	{"aabcdaaaaaaaaaa", "a2b1c1d1a10"},
	{"aabcdeaaaaaaaaaa", "a2b1c1d1e1a10"},
	{"aabcdefaaaaaaaaaa", "a2b1c1d1e1f1a10"},
	{"aabcdefgaaaaaaaaaa", "a2b1c1d1e1f1g1a10"},
	{"aabcdefghaaaaaaaaaa", "aabcdefghaaaaaaaaaa"},
	{
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
		"a100b100",
	},
	{
		"aabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabb",
		"aabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabb",
	},
	{
		"aaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbbaaabbb",
		"a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3a3b3",
	},
	{
		"aaaaaaaaaabbbbbbbbbbaaaaaaaaaabbbbbbbbbbaaaaaaaaaabbbbbbbbbbaaaaaaaaaabbbbbbbbbbaaaaaaaaaabbbbbbbbbbaaaaaaaaaabbbbbbbbbbaaaaaaaaaabbbbbbbbbbaaaaaaaaaabbbbbbbbbbaaaaaaaaaabbbbbbbbbbaaaaaaaaaabbbbbbbbbbaaaaaaaaaabbbbbbbbbb",
		"a10b10a10b10a10b10a10b10a10b10a10b10a10b10a10b10a10b10a10b10a10b10",
	},
}

func TestCompress(t *testing.T) {
	for _, c := range cases {
		actual := compress(c.input)
		if actual != c.expected {
			t.Errorf("input: %q want: %q got: %q", c.input, c.expected, actual)
		}
	}
}

func TestCompressCountFirst(t *testing.T) {
	for _, c := range cases {
		actual := compressCountFirst(c.input)
		if actual != c.expected {
			t.Errorf("input: %q want: %q got: %q", c.input, c.expected, actual)
		}
	}
}

func BenchmarkCompress(b *testing.B) {
	for _, c := range cases {
		name := c.expected
		if len(name) == 0 {
			name = "empty"
		}
		if len(name) > 20 {
			name = name[:17] + "..."
		}
		b.Run(name, func(b2 *testing.B) {
			b2.ReportAllocs()
			for i := 0; i < b2.N; i++ {
				compress(c.input)
			}
		})
	}
}

func BenchmarkCompressCountFirst(b *testing.B) {
	for _, c := range cases {
		name := c.expected
		if len(name) == 0 {
			name = "empty"
		}
		if len(name) > 20 {
			name = name[:17] + "..."
		}
		b.Run(name, func(b2 *testing.B) {
			b2.ReportAllocs()
			for i := 0; i < b2.N; i++ {
				compressCountFirst(c.input)
			}
		})
	}
}
