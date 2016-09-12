package zeromatrix

import (
	"bytes"
	"fmt"
	"testing"
)

type testData struct {
	input    [][]int
	expected [][]int
}

func TestZeroUsingMap(t *testing.T) {
	for _, c := range getCases() {
		zeroUsingMap(c.input)
		if !areEqual(c.input, c.expected) {
			t.Errorf("\nwant:\n%s\ngot:\n%s\n",
				getMatrixString(c.expected),
				getMatrixString(c.input))
		}
	}
}

func TestZeroUsingSlice(t *testing.T) {
	for _, c := range getCases() {
		zeroUsingSlice(c.input)
		if !areEqual(c.input, c.expected) {
			t.Errorf("\nwant:\n%s\ngot:\n%s\n",
				getMatrixString(c.expected),
				getMatrixString(c.input))
		}
	}
}

func BenchmarkZeroUsingMap(b *testing.B) {
	benchmark(b, zeroUsingMap)
}

func BenchmarkZeroUsingSlice(b *testing.B) {
	benchmark(b, zeroUsingSlice)
}

func benchmark(b *testing.B, zeroFunc func([][]int)) {
	for _, c := range getCases() {
		var name string
		if len(c.input) == 0 {
			name = "0x0"
		} else {
			name = fmt.Sprintf("%dx%d", len(c.input), len(c.input[0]))
		}
		b.Run(name, func(b2 *testing.B) {
			m := make([][]int, len(c.input))
			for i := 0; i < len(c.input); i++ {
				m[i] = make([]int, len(c.input[i]))
			}
			b2.ResetTimer()
			b2.ReportAllocs()
			for i := 0; i < b2.N; i++ {
				copyMatrix(m, c.input)
				zeroFunc(m)
			}
		})
	}
}

func copyMatrix(dst, src [][]int) {
	for r := 0; r < len(src); r++ {
		copy(dst[r], src[r])
	}
}

func areEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		x, y := a[i], b[i]
		if len(x) != len(y) {
			return false
		}
		for j := 0; j < len(x); j++ {
			if x[j] != y[j] {
				return false
			}
		}
	}
	return true
}

func getMatrixString(m [][]int) string {
	b := bytes.Buffer{}
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			b.WriteString(fmt.Sprintf("%d ", m[y][x]))
		}
		b.WriteString("\n")
	}
	return b.String()
}

func getCases() []testData {
	return []testData{
		{
			input:    [][]int{},
			expected: [][]int{},
		},
		{
			input: [][]int{
				{0},
			},
			expected: [][]int{
				{0},
			},
		},
		{
			input: [][]int{
				{1},
			},
			expected: [][]int{
				{1},
			},
		},
		{
			input: [][]int{
				{1, 0, 0},
				{1, 1, 1},
				{1, 1, 1},
			},
			expected: [][]int{
				{0, 0, 0},
				{1, 0, 0},
				{1, 0, 0},
			},
		},
		{
			input: [][]int{
				{1, 0},
				{1, 1},
				{1, 1},
			},
			expected: [][]int{
				{0, 0},
				{1, 0},
				{1, 0},
			},
		},
		{
			input: [][]int{
				{1, 0, 0},
				{1, 1, 1},
			},
			expected: [][]int{
				{0, 0, 0},
				{1, 0, 0},
			},
		},
		{
			input: [][]int{
				{0},
				{1},
				{1},
			},
			expected: [][]int{
				{0},
				{0},
				{0},
			},
		},
		{
			input: [][]int{
				{0, 1, 1},
			},
			expected: [][]int{
				{0, 0, 0},
			},
		},
		{
			input: [][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 1, 1},
			},
			expected: [][]int{
				{1, 1, 0},
				{0, 0, 0},
				{1, 1, 0},
			},
		},
		{
			input: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			expected: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
		},
		{
			input: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			expected: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			input: [][]int{
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1},
			},
			expected: [][]int{
				{1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 0, 1, 0, 0, 1, 1},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 0, 1, 0, 0, 1, 1},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 0, 1, 0, 0, 1, 1},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	}
}
