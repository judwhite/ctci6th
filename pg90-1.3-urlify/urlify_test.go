package urlify

import "testing"

// URLify: Write a method to replace all spaces in a string with '%20'. You may assume that the string has
// sufficient space at the end to hold the additional characters, and that you are given the "true" length of
// the string. (Note: If implementing in Java, please use a character array so that you can perform this operation
// in place.)
//
// EXAMPLE
// Input: "Mr John Smith    ", 13
// Output: "Mr%20John%20Smith"

var cases = []struct {
	input      string
	trueLength int
	expected   string
}{
	{"Mr John Smith    ", 13, "Mr%20John%20Smith"},
	{"Mr John A. Smith      ", 16, "Mr%20John%20A.%20Smith"},
	{"Mr John 界. Smith      ", 18, "Mr%20John%20界.%20Smith"}, // len("界") = 3
	{"Mr John A. Smith         ", 17, "Mr%20John%20A.%20Smith%20"},
	{" Mr John A. Smith           ", 18, "%20Mr%20John%20A.%20Smith%20"},
	{"  Mr John A. Smith             ", 19, "%20%20Mr%20John%20A.%20Smith%20"},
	{"  Mr John A.  Smith               ", 20, "%20%20Mr%20John%20A.%20%20Smith%20"},
	{"  Mr John A.  Smith abcdefghijklmnopqrstuvwxyz0123456789              ",
		56,
		"%20%20Mr%20John%20A.%20%20Smith%20abcdefghijklmnopqrstuvwxyz0123456789",
	},
}

func TestURLifyInPlace(t *testing.T) {
	for _, c := range cases {
		actual := urlifyInPlace(c.input, c.trueLength)
		if actual != c.expected {
			t.Errorf("%q want: %q got: %q", c.input, c.expected, actual)
		}
	}
}

func TestURLifyReplace(t *testing.T) {
	for _, c := range cases {
		actual := urlifyReplace(c.input, c.trueLength)
		if actual != c.expected {
			t.Errorf("%q want: %q got: %q", c.input, c.expected, actual)
		}
	}
}

func TestURLifyAllocNew(t *testing.T) {
	for _, c := range cases {
		actual := urlifyAllocNew(c.input, c.trueLength)
		if actual != c.expected {
			t.Errorf("%q want: %q got: %q", c.input, c.expected, actual)
		}
	}
}

func TestURLifyInPlace2(t *testing.T) {
	for _, c := range cases {
		actual := urlifyInPlace2(c.input, c.trueLength)
		if actual != c.expected {
			t.Errorf("%q want: %q got: %q", c.input, c.expected, actual)
		}
	}
}

func TestURLifyInPlaceBytes(t *testing.T) {
	for _, c := range cases {
		actual := urlifyInPlaceBytes([]byte(c.input), c.trueLength)
		if string(actual) != c.expected {
			t.Errorf("%q want: %q got: %q", c.input, c.expected, string(actual))
		}
	}
}

func TestURLifyInPlaceBytes2(t *testing.T) {
	for _, c := range cases {
		actual := urlifyInPlaceBytes2([]byte(c.input), c.trueLength)
		if string(actual) != c.expected {
			t.Errorf("%q want: %q got: %q", c.input, c.expected, string(actual))
		}
	}
}

const benchmarkStringInput string = "  Mr John 界.  Smith               "
const benchmarkStringLength = 22

var benchmarkByteInput []byte = []byte(benchmarkStringInput)

func BenchmarkURLifyInPlace(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		urlifyInPlace(benchmarkStringInput, benchmarkStringLength)
	}
}

func BenchmarkURLifyReplace(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		urlifyReplace(benchmarkStringInput, benchmarkStringLength)
	}
}

func BenchmarkURLifyAllocNew(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		urlifyAllocNew(benchmarkStringInput, benchmarkStringLength)
	}
}

func BenchmarkURLifyInPlace2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		urlifyInPlace2(benchmarkStringInput, benchmarkStringLength)
	}
}

func BenchmarkURLifyInPlaceBytes(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		urlifyInPlaceBytes(benchmarkByteInput, benchmarkStringLength)
	}
}

func BenchmarkURLifyInPlaceBytes2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		urlifyInPlaceBytes2(benchmarkByteInput, benchmarkStringLength)
	}
}
