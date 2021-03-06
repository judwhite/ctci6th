package pairs

import (
	"math/rand"
	"testing"
	"time"
)

// Given an array of distinct integer values, count the number of pairs of integers that have difference k.
// For example, given the array {1, 7, 5, 9, 2, 12, 3} and the difference k = 2, there are four pairs with
// difference 2: (1, 3), (3, 5), (5, 7), (7, 9)

func TestBruteForce(t *testing.T) {
	input := []int{1, 7, 5, 9, 2, 12, 3}
	k := 2
	expected := [][]int{{1, 3}, {3, 5}, {5, 7}, {7, 9}}

	actual := bruteForce(input, k)

	validate(t, "bruteForce", expected, actual)
}

func TestSortFirst(t *testing.T) {
	input := []int{1, 7, 5, 9, 2, 12, 3}
	k := 2
	expected := [][]int{{1, 3}, {3, 5}, {5, 7}, {7, 9}}

	actual := sortFirst(input, k)

	validate(t, "sortFirst", expected, actual)
}

func TestHashMap(t *testing.T) {
	input := []int{1, 7, 5, 9, 2, 12, 3}
	k := 2
	expected := [][]int{{1, 3}, {3, 5}, {5, 7}, {7, 9}}

	actual := hashMap(input, k)

	validate(t, "hashMap", expected, actual)
}

func TestSparseSet(t *testing.T) {
	input := []int{1, 7, 5, 9, 2, 12, 3}
	k := 2
	expected := [][]int{{1, 3}, {3, 5}, {5, 7}, {7, 9}}

	actual := sparseSet(input, k)

	validate(t, "sparseSet", expected, actual)
}

type testCase struct {
	name string
	f    func([]int, int) [][]int
}

func TestAllAgainstBruteForce(t *testing.T) {
	input := getRandomInput()
	k := 6

	expected := bruteForce(input, k)

	cases := []testCase{
		{"sortFirst", sortFirst},
		{"hashMap", hashMap},
		{"sparseSet", sparseSet},
	}
	for _, c := range cases {
		actual := c.f(input, k)
		validate(t, c.name, expected, actual)
	}
}

func BenchmarkBruteForce(b *testing.B) {
	var input []int
	for i := int(1e5); i >= 0; i -= 3 {
		input = append(input, i)
	}
	k := 6

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bruteForce(input, k)
	}
}

func BenchmarkSortFirst(b *testing.B) {
	var input []int
	for i := int(1e5); i >= 0; i -= 3 {
		input = append(input, i)
	}
	k := 6

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortFirst(input, k)
	}
}

func BenchmarkHashMap(b *testing.B) {
	var input []int
	for i := int(1e5); i >= 0; i -= 3 {
		input = append(input, i)
	}
	k := 6

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hashMap(input, k)
	}
}

func BenchmarkSparseSet(b *testing.B) {
	var input []int
	for i := int(1e5); i >= 0; i -= 3 {
		input = append(input, i)
	}
	k := 6

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sparseSet(input, k)
	}
}

var randomInput []int

func getRandomInput() []int {
	if len(randomInput) == 0 {
		rand.Seed(time.Now().UnixNano())
		set := make(map[int]bool)
		const size = 1e5
		for i := 0; i < size; i++ {
			n := rand.Intn(size * 3)
			for set[n] {
				n = rand.Intn(size * 3)
			}
			set[n] = true
			randomInput = append(randomInput, n)
		}
	}

	return randomInput
}

func BenchmarkBruteForceRandom(b *testing.B) {
	input := getRandomInput()
	k := 6

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bruteForce(input, k)
	}
}

func BenchmarkSortFirstRandom(b *testing.B) {
	input := getRandomInput()
	k := 6

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortFirst(input, k)
	}
}

func BenchmarkHashMapRandom(b *testing.B) {
	input := getRandomInput()
	k := 6

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hashMap(input, k)
	}
}

func BenchmarkSparseSetRandom(b *testing.B) {
	input := getRandomInput()
	k := 6

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sparseSet(input, k)
	}
}

func validate(t testing.TB, name string, expected, actual [][]int) {
	if len(expected) != len(actual) {
		t.Errorf("func: %s\n\t\t  want: %v\n\t\t  got:  %v\n\t\t  len(want) = %d len(got) = %d",
			name, expected, actual, len(expected), len(actual))
		return
	}

	marked := make([]bool, len(actual))
	for i := 0; i < len(expected); i++ {
		// order of pairs doesn't matter
		found := false
		for j := 0; j < len(actual); j++ {
			if marked[j] {
				continue
			}
			if equal(expected[i], actual[j]) {
				marked[j] = true
				found = true
				break
			}
		}
		if !found {
			t.Errorf("want: %v got: %v", expected, actual)
			return
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
