package sumlists

import (
	"fmt"
	"strconv"
	"testing"
)

var cases = []struct {
	a   []int
	b   []int
	sum []int
}{
	{a: []int{0}, b: []int{0}, sum: []int{0}},
	{a: []int{1}, b: []int{0}, sum: []int{1}},
	{a: []int{0}, b: []int{1}, sum: []int{1}},
	{a: []int{1}, b: []int{1}, sum: []int{2}},
	{a: []int{9}, b: []int{1}, sum: []int{1, 0}},
	{a: []int{1}, b: []int{9}, sum: []int{1, 0}},
	{a: []int{1, 0}, b: []int{0}, sum: []int{1, 0}},
	{a: []int{0}, b: []int{1, 0}, sum: []int{1, 0}},
	{a: []int{9, 9, 9}, b: []int{1}, sum: []int{1, 0, 0, 0}},
	{a: []int{1}, b: []int{9, 9, 9}, sum: []int{1, 0, 0, 0}},
	{a: []int{6, 1, 7}, b: []int{2, 9, 5}, sum: []int{9, 1, 2}},
	{a: []int{8, 2, 4, 5}, b: []int{5, 9, 3, 5}, sum: []int{1, 4, 1, 8, 0}},
	{a: []int{9, 1, 4, 9}, b: []int{7, 0}, sum: []int{9, 2, 1, 9}},
	{a: []int{9, 0, 0}, b: []int{3, 8, 3, 9}, sum: []int{4, 7, 3, 9}},
	{a: []int{3, 3, 1}, b: []int{8, 9, 3, 2}, sum: []int{9, 2, 6, 3}},
	{a: []int{9, 6, 3, 5}, b: []int{8, 6, 8, 7}, sum: []int{1, 8, 3, 2, 2}},
	{a: []int{6, 4}, b: []int{3, 8, 0, 1}, sum: []int{3, 8, 6, 5}},
	{a: []int{2, 1, 5, 4}, b: []int{3, 3, 5, 2}, sum: []int{5, 5, 0, 6}},
	{a: []int{6, 4, 3, 3}, b: []int{2, 5, 9, 1}, sum: []int{9, 0, 2, 4}},
	{a: []int{4, 3, 6}, b: []int{8, 5, 0, 9}, sum: []int{8, 9, 4, 5}},
	{a: []int{7, 4, 2, 0}, b: []int{2, 1, 7, 7}, sum: []int{9, 5, 9, 7}},
	{a: []int{6, 8, 6, 6}, b: []int{7, 9, 7, 8}, sum: []int{1, 4, 8, 4, 4}},
	{a: []int{8, 9, 9, 0}, b: []int{1, 0, 5, 8}, sum: []int{1, 0, 0, 4, 8}},
	{a: []int{5, 1, 0, 8}, b: []int{6, 1, 5, 6}, sum: []int{1, 1, 2, 6, 4}},
	{a: []int{6, 6, 9, 0}, b: []int{8, 5, 3, 0}, sum: []int{1, 5, 2, 2, 0}},
	{a: []int{6, 9, 8, 0}, b: []int{3, 4, 1, 1}, sum: []int{1, 0, 3, 9, 1}},
	{a: []int{3, 0, 6, 3}, b: []int{4, 3, 3, 1}, sum: []int{7, 3, 9, 4}},
	{a: []int{3, 3, 9, 0}, b: []int{6, 4, 3, 8}, sum: []int{9, 8, 2, 8}},
	{a: []int{6, 0, 1, 9}, b: []int{8, 2, 9, 6}, sum: []int{1, 4, 3, 1, 5}},
	{a: []int{4, 0, 1, 6}, b: []int{7, 9, 3, 7}, sum: []int{1, 1, 9, 5, 3}},
	{a: []int{2, 8, 8, 9}, b: []int{2, 8, 1}, sum: []int{3, 1, 7, 0}},
	{a: []int{2, 0, 0, 0}, b: []int{2, 0, 0}, sum: []int{2, 2, 0, 0}},
}

// Input: (7 -> 1 -> 6) + (5 -> 9 -> 2). That is, 617 + 295.
// Output: 2 -> 1 -> 9. That is, 912.

func TestSum(t *testing.T) {
	for _, c := range cases {
		a := toLinkedList(c.a)
		b := toLinkedList(c.b)

		actual := sum(a, b)

		actualSlice := toSlice(actual)

		if !equal(actualSlice, c.sum) {
			t.Errorf("%v + %v want: %v got: %v", c.a, c.b, c.sum, actualSlice)
		}
	}
}

func TestSumReverse(t *testing.T) {
	for _, c := range cases {
		ra := reverse(c.a)
		rb := reverse(c.b)
		rsum := reverse(c.sum)

		a := toLinkedList(ra)
		b := toLinkedList(rb)

		actual := sumReverse(a, b)

		actualSlice := toSlice(actual)

		if !equal(actualSlice, rsum) {
			t.Errorf("%v + %v want: %v got: %v", ra, rb, rsum, actualSlice)
		}
	}
}

func TestSumBook(t *testing.T) {
	for _, c := range cases {
		a := toLinkedList(c.a)
		b := toLinkedList(c.b)

		actual := sumBook(a, b)

		actualSlice := toSlice(actual)

		if !equal(actualSlice, c.sum) {
			t.Errorf("%v + %v want: %v got: %v", c.a, c.b, c.sum, actualSlice)
		}
	}
}

func TestSumReverseBook(t *testing.T) {
	for _, c := range cases {
		ra := reverse(c.a)
		rb := reverse(c.b)
		rsum := reverse(c.sum)

		a := toLinkedList(ra)
		b := toLinkedList(rb)

		actual := sumReverseBook(a, b)

		actualSlice := toSlice(actual)

		if !equal(actualSlice, rsum) {
			t.Errorf("%v + %v want: %v got: %v", ra, rb, rsum, actualSlice)
		}
	}
}

func BenchmarkSum(b *testing.B) {
	for _, c := range cases {
		n1 := toLinkedList(c.a)
		n2 := toLinkedList(c.b)
		name := fmt.Sprintf("%s+%s=%s", toString(c.a), toString(c.b), toString(c.sum))
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				sum(n1, n2)
			}
		})
	}
}

func BenchmarkSumReverse(b *testing.B) {
	for _, c := range cases {
		n1 := toLinkedList(reverse(c.a))
		n2 := toLinkedList(reverse(c.b))
		name := fmt.Sprintf("%s+%s=%s", toString(c.a), toString(c.b), toString(c.sum))
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				sumReverse(n1, n2)
			}
		})
	}
}

func BenchmarkSumBook(b *testing.B) {
	for _, c := range cases {
		n1 := toLinkedList(c.a)
		n2 := toLinkedList(c.b)
		name := fmt.Sprintf("%s+%s=%s", toString(c.a), toString(c.b), toString(c.sum))
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				sumBook(n1, n2)
			}
		})
	}
}

func BenchmarkSumReverseBook(b *testing.B) {
	for _, c := range cases {
		n1 := toLinkedList(reverse(c.a))
		n2 := toLinkedList(reverse(c.b))
		name := fmt.Sprintf("%s+%s=%s", toString(c.a), toString(c.b), toString(c.sum))
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				sumReverseBook(n1, n2)
			}
		})
	}
}

func toString(a []int) string {
	var s string
	for _, v := range a {
		s += strconv.Itoa(v)
	}
	return s
}

func reverse(a []int) []int {
	b := make([]int, len(a))
	for i, j := 0, len(a)-1; j >= 0; {
		b[i] = a[j]
		i++
		j--
	}
	return b
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

func toLinkedList(list []int) *node {
	if len(list) == 0 {
		return nil
	}
	head := &node{val: list[0]}
	for i, n := 1, head; i < len(list); i++ {
		n.next = &node{val: list[i]}
		n = n.next
	}
	return head
}

func toSlice(head *node) []int {
	if head == nil {
		return nil
	}
	var list []int
	for n := head; n != nil; n = n.next {
		list = append(list, n.val)
	}
	return list
}
