package pairs

import (
	"sort"

	"golang.org/x/tools/container/intsets"
)

// Given an array of distinct integer values, count the number of pairs of integers that have difference k.
// For example, given the array {1, 7, 5, 9, 2, 12, 3} and the difference k = 2, there are four pairs with
// difference 2: (1, 3), (3, 5), (5, 7), (7, 9)

func bruteForce(list []int, k int) [][]int {
	var result [][]int

	if k < 0 {
		k = -k
	}
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			a, b := order(list[i], list[j])
			if b-a == k {
				result = append(result, []int{a, b})
			}
		}
	}

	return result
}

func sortFirst(list []int, k int) [][]int {
	var result [][]int

	sort.Ints(list)
	if k < 0 {
		k = -k
	}
	for i := 0; i < len(list)-1; i++ {
		a := list[i]
		for j := i + 1; j < len(list); j++ {
			b := list[j]
			d := b - a
			if d > k {
				break
			}
			if d == k {
				result = append(result, []int{a, b})
			}
		}
	}
	return result
}

func hashMap(list []int, k int) [][]int {
	var result [][]int
	test := make(map[int]struct{}, len(list))

	for _, val := range list {
		a, b := val-k, val+k
		if _, ok := test[a]; ok {
			x, y := order(a, val)
			result = append(result, []int{x, y})
		}
		if _, ok := test[b]; ok {
			x, y := order(b, val)
			result = append(result, []int{x, y})
		}
		test[val] = struct{}{}
	}

	return result
}

func sparseSet(list []int, k int) [][]int {
	var result [][]int
	set := intsets.Sparse{}

	for _, val := range list {
		a, b := val-k, val+k
		if set.Has(a) {
			x, y := order(a, val)
			result = append(result, []int{x, y})
		}
		if set.Has(b) {
			x, y := order(b, val)
			result = append(result, []int{x, y})
		}
		set.Insert(val)
	}

	return result
}

func order(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}
