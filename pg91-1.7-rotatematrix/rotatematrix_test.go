package rotatematrix

import (
	"bytes"
	"fmt"
	"testing"
)

type testCase struct {
	input    [][]uint32
	expected [][]uint32
}

func TestRotate(t *testing.T) {
	for _, c := range getCases() {
		rotate(c.input)
		if !areEqual(c.input, c.expected) {
			t.Errorf("\nwant:\n%s\ngot:\n%s\n",
				getMatrixString(c.expected),
				getMatrixString(c.input))
		}
	}
}

func BenchmarkRotate(b *testing.B) {
	for _, c := range getCases() {
		b.Run(fmt.Sprintf("%dx%d", len(c.input), len(c.input)), func(b2 *testing.B) {
			b2.ReportAllocs()
			for i := 0; i < b2.N; i++ {
				rotate(c.input)
			}
		})
	}
}

func areEqual(a, b [][]uint32) bool {
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

func getMatrixString(m [][]uint32) string {
	b := bytes.Buffer{}
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			b.WriteString(fmt.Sprintf("%02d ", m[y][x]))
		}
		b.WriteString("\n")
	}
	return b.String()
}

func getCases() []testCase {
	return []testCase{
		{
			input:    [][]uint32{},
			expected: [][]uint32{},
		},
		{
			input:    [][]uint32{{0}},
			expected: [][]uint32{{0}},
		},
		{
			input: [][]uint32{
				{0, 1},
				{2, 3},
			},
			expected: [][]uint32{
				{2, 0},
				{3, 1},
			},
		},
		{
			input: [][]uint32{
				{0, 1, 2},
				{3, 4, 5},
				{6, 7, 8},
			},
			expected: [][]uint32{
				{6, 3, 0},
				{7, 4, 1},
				{8, 5, 2},
			},
		},
		{
			input: [][]uint32{
				{0, 1, 2, 3},
				{4, 5, 6, 7},
				{8, 9, 10, 11},
				{12, 13, 14, 15},
			},
			expected: [][]uint32{
				{12, 8, 4, 0},
				{13, 9, 5, 1},
				{14, 10, 6, 2},
				{15, 11, 7, 3},
			},
		},
		{
			input: [][]uint32{
				{0, 1, 2, 3, 4},
				{5, 6, 7, 8, 9},
				{10, 11, 12, 13, 14},
				{15, 16, 17, 18, 19},
				{20, 21, 22, 23, 24},
			},
			expected: [][]uint32{
				{20, 15, 10, 5, 0},
				{21, 16, 11, 6, 1},
				{22, 17, 12, 7, 2},
				{23, 18, 13, 8, 3},
				{24, 19, 14, 9, 4},
			},
		},
		{
			input: [][]uint32{
				{0, 1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10, 11},
				{12, 13, 14, 15, 16, 17},
				{18, 19, 20, 21, 22, 23},
				{24, 25, 26, 27, 28, 29},
				{30, 31, 32, 33, 34, 35},
			},
			expected: [][]uint32{
				{30, 24, 18, 12, 6, 0},
				{31, 25, 19, 13, 7, 1},
				{32, 26, 20, 14, 8, 2},
				{33, 27, 21, 15, 9, 3},
				{34, 28, 22, 16, 10, 4},
				{35, 29, 23, 17, 11, 5},
			},
		},
	}
}
