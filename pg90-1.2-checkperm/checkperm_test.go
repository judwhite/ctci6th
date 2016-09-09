package checkperm

import (
	"testing"
)

func TestIsPermMap(t *testing.T) {
	cases := []struct {
		a        string
		b        string
		expected bool
	}{
		{"", "", true},
		{"abcd", "abdc", true},
		{"abcd", "abcc", false},
		{"smaismrmilmepoetaleumibunenvgttaviras", "altissimumplanetamtergeminumobservavi", true},
		{"mrmojorisin", "jimmorrison", true},
		{"rescued", "secured", true},
	}

	for _, c := range cases {
		actual := isPermMap(c.a, c.b)
		if actual != c.expected {
			t.Errorf("a: %q b: %q want: %v got: %v", c.a, c.b, c.expected, actual)
		}
	}
}

func TestIsPermArray(t *testing.T) {
	cases := []struct {
		a        string
		b        string
		expected bool
	}{
		{"", "", true},
		{"abcd", "abdc", true},
		{"abcd", "abcc", false},
		{"smaismrmilmepoetaleumibunenvgttaviras", "altissimumplanetamtergeminumobservavi", true},
		{"mrmojorisin", "jimmorrison", true},
		{"rescued", "secured", true},
	}

	for _, c := range cases {
		actual := isPermArray(c.a, c.b)
		if actual != c.expected {
			t.Errorf("a: %q b: %q want: %v got: %v", c.a, c.b, c.expected, actual)
		}
	}
}

func TestIsPermSort(t *testing.T) {
	cases := []struct {
		a        string
		b        string
		expected bool
	}{
		{"", "", true},
		{"abcd", "abdc", true},
		{"abcd", "abcc", false},
		{"smaismrmilmepoetaleumibunenvgttaviras", "altissimumplanetamtergeminumobservavi", true},
		{"mrmojorisin", "jimmorrison", true},
		{"rescued", "secured", true},
	}

	for _, c := range cases {
		actual := isPermSort(c.a, c.b)
		if actual != c.expected {
			t.Errorf("a: %q b: %q want: %v got: %v", c.a, c.b, c.expected, actual)
		}
	}
}

func BenchmarkIsPermMapLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPermMap(
			"smaismrmilmepoetaleumibunenvgttavirassmaismrmilmepoetaleumibunenvgttaviras",
			"altissimumplanetamtergeminumobservavialtissimumplanetamtergeminumobservavi")
	}
}

func BenchmarkIsPermArrayLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPermArray(
			"smaismrmilmepoetaleumibunenvgttavirassmaismrmilmepoetaleumibunenvgttaviras",
			"altissimumplanetamtergeminumobservavialtissimumplanetamtergeminumobservavi")
	}
}

func BenchmarkIsPermSortLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPermSort(
			"smaismrmilmepoetaleumibunenvgttavirassmaismrmilmepoetaleumibunenvgttaviras",
			"altissimumplanetamtergeminumobservavialtissimumplanetamtergeminumobservavi")
	}
}

func BenchmarkIsPermMapShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPermMap("abc", "cba")
	}
}

func BenchmarkIsPermArrayShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPermArray("abc", "cba")
	}
}

func BenchmarkIsPermSortShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPermSort("abc", "cba")
	}
}
