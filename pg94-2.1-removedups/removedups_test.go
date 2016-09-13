package removedups

import (
	"fmt"
	"testing"
)

var cases = []struct {
	list     []int
	expected []int
}{
	{
		list:     []int{},
		expected: []int{},
	},
	{
		list:     []int{1},
		expected: []int{1},
	},
	{
		list:     []int{1, 1},
		expected: []int{1},
	},
	{
		list:     []int{1, 2, 1},
		expected: []int{1, 2},
	},
	{
		list:     []int{2, 1, 1},
		expected: []int{2, 1},
	},
	{
		list:     []int{2, 2, 2, 2},
		expected: []int{2},
	},
	{
		list:     []int{2, 3, 2, 4, 2, 5, 2, 6},
		expected: []int{2, 3, 4, 5, 6},
	},
	{
		list:     []int{1, 2, 3, 4, 2, 5, 6, 2, 3, 1, 5, 6, 9, 0, 1},
		expected: []int{1, 2, 3, 4, 5, 6, 9, 0},
	},
	{
		list:     []int{-1, 1, 2, 3, 4, 2, 5, 6, 2, 3, 1, 5, 6, 9, 0, 1},
		expected: []int{-1, 1, 2, 3, 4, 5, 6, 9, 0},
	},
	{
		list: []int{
			-1, 1, 2, 3, 4, 2, 5, 6, 2, 3, 1, 5, 6, 9, 0, 1,
			-1, 1, 2, 3, 4, 2, 5, 6, 2, 3, 1, 5, 6, 9, 0, 1,
			-1, 1, 2, 3, 4, 2, 5, 6, 2, 3, 1, 5, 6, 9, 0, 1,
			-1, 1, 2, 3, 4, 2, 5, 6, 2, 3, 1, 5, 6, 9, 0, 1,
			-1, 1, 2, 3, 4, 2, 5, 6, 2, 3, 1, 5, 6, 9, 0, 1,
			-1, 1, 2, 3, 4, 2, 5, 6, 2, 3, 1, 5, 6, 9, 0, 1,
			-1, 1, 2, 3, 4, 2, 5, 6, 2, 3, 1, 5, 6, 9, 0, 1,
			-1, 1, 2, 3, 4, 2, 5, 6, 2, 3, 1, 5, 6, 9, 0, 1,
			-1, 1, 2, 3, 4, 2, 5, 6, 2, 3, 1, 5, 6, 9, 0, 1,
		},
		expected: []int{-1, 1, 2, 3, 4, 5, 6, 9, 0},
	},
	{
		list: []int{
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
			10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
			20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
			30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
			40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
			50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
			60, 61, 62, 63, 64, 65, 66, 67, 68, 69,
			70, 71, 72, 73, 74, 75, 76, 77, 78, 79,
			80, 81, 82, 83, 84, 85, 86, 87, 88, 89,
			90, 91, 92, 93, 94, 95, 96, 97, 98, 99,
		},
		expected: []int{
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
			10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
			20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
			30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
			40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
			50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
			60, 61, 62, 63, 64, 65, 66, 67, 68, 69,
			70, 71, 72, 73, 74, 75, 76, 77, 78, 79,
			80, 81, 82, 83, 84, 85, 86, 87, 88, 89,
			90, 91, 92, 93, 94, 95, 96, 97, 98, 99,
		},
	},
}

func TestRemoveDupes(t *testing.T) {
	for _, c := range cases {
		head := toLinkedList(c.list)
		removeDupes(head)
		actual := toSlice(head)
		if !equal(c.expected, actual) {
			t.Errorf("list: %v want: %v got: %v", c.list, c.expected, actual)
		}
	}
}

func TestRemoveDupesNoBuffer(t *testing.T) {
	for _, c := range cases {
		head := toLinkedList(c.list)
		removeDupesNoBuffer(head)
		actual := toSlice(head)
		if !equal(c.expected, actual) {
			t.Errorf("list: %v want: %v got: %v", c.list, c.expected, actual)
		}
	}
}

func BenchmarkRemoveDupes(b *testing.B) {
	for _, c := range cases {
		b.Run(fmt.Sprintf("len%d_ddlen%d", len(c.list), len(c.expected)), func(b *testing.B) {
			heads := make([]*node, b.N)
			for i := 0; i < b.N; i++ {
				heads[i] = toLinkedList(c.list)
			}
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				removeDupes(heads[i])
			}
		})
	}
}

func BenchmarkRemoveDupesNoBuffer(b *testing.B) {
	for _, c := range cases {
		b.Run(fmt.Sprintf("len%d_ddlen%d", len(c.list), len(c.expected)), func(b *testing.B) {
			heads := make([]*node, b.N)
			for i := 0; i < b.N; i++ {
				heads[i] = toLinkedList(c.list)
			}
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				removeDupesNoBuffer(heads[i])
			}
		})
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
