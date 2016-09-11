package isunique

import "testing"

var cases = []struct {
	input    string
	isASCII  bool
	isBytes  bool
	expected bool
}{
	{"", true, false, true},
	{"a", true, false, true},
	{"abc", true, false, true},
	{"baeac", true, false, false},
	{"zazb", true, false, false},
	{"bzaz", true, false, false},
	{"zbza", true, false, false},
	{"zzba", true, false, false},
	{"bazz", true, false, false},
	{"zyxwvutsrqponmlkjihgfedcba", true, false, true},
	{"zyxwvutsrqponmlkjihgfedcbaz", true, false, false},
	{"zyxwvutsrqponmlkjihgfedcbaa", false, false, false},
	{"azyxwvutsrqponmlkjihgfedcba", false, false, false},
	{"Helo, 世界", false, false, true},
	{"界世界", false, false, false},
	{"\x00\xFF\x40\x7f\x80\xc0\xbf\xfe\x01\x81\xc1\x41", true, true, true},
	{"\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98", true, true, true},
	{"\xbd\xe2\xb2\x3d\xbc\x20\xe2\x8c\x98", true, true, false},
}

func TestIsUniqueSort(t *testing.T) {
	for _, c := range cases {
		if c.isBytes {
			continue
		}
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
		if c.isBytes {
			continue
		}
		actual := isUniqueMap(c.input)
		if actual != c.expected {
			t.Errorf("%q want: %v got: %v", c.input, c.expected, actual)
		}
	}
}

func TestIsUniqueBitVector(t *testing.T) {
	for _, c := range cases {
		if !c.isASCII {
			// TestIsUniqueBitVector only supports ASCII characters
			continue
		}
		actual := isUniqueBitVector(c.input)
		if actual != c.expected {
			t.Errorf("%q want: %v got: %v", c.input, c.expected, actual)
		}
	}

	// special cases just for bit vector
	b := make([]byte, 256)
	for i := 0; i < 256; i++ {
		b[i] = byte(i)
	}

	s := string(b)
	actual := isUniqueBitVector(s)
	if !actual {
		t.Errorf("%q want: %v got: %v", s, true, actual)
	}

	b = append(b, 0xFF)
	s = string(b)
	actual = isUniqueBitVector(s)
	if actual {
		t.Errorf("%q want: %v got: %v", s, false, actual)
	}
}

func TestIsUniqueBitVector2(t *testing.T) {
	for _, c := range cases {
		if !c.isASCII {
			// TestIsUniqueBitVector2 only supports a-z ASCII characters
			continue
		}
		// TestIsUniqueBitVector2 only supports a-z ASCII characters
		ok := true
		for i := 0; i < len(c.input); i++ {
			if c.input[i] < 'a' || c.input[i] > 'z' {
				ok = false
				break
			}
		}
		if !ok {
			continue
		}

		actual := isUniqueBitVector(c.input)
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

func BenchmarkIsUniqueBitVector(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isUniqueBitVector("zyxwvutsrqponmlkjihgfedcba")
	}
}

func BenchmarkIsUniqueBitVector2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isUniqueBitVector2("zyxwvutsrqponmlkjihgfedcba")
	}
}
