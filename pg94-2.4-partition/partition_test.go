package partition

import (
	"errors"
	"fmt"
	"sort"
	"testing"
)

var cases = []struct {
	list []int
	k    int
}{
	{list: []int{3, 5, 8, 5, 10, 2, 1}, k: 5},
	{list: []int{3, 5, 8, 5, 10, 2, 1, 1}, k: 5},
	{list: []int{1, 1, 1, 1, 1}, k: 5},
	{list: []int{5, 5, 5}, k: 5},
	{list: []int{10, 5, 4}, k: 5},
	{list: []int{10, 5, 4, 3, 2, 1}, k: 5},
	{list: []int{2, 1, 5, 5, 9, 9, 0, 3}, k: 9},
	{list: []int{9, 1, 1}, k: 9},
	{list: []int{9, 9, 1}, k: 9},
	{list: []int{1, 1, 9}, k: 9},
	{list: []int{1, 9, 9}, k: 9},
	{list: []int{1, 9, 1}, k: 9},
	{list: []int{9, 1, 9}, k: 9},
	{
		list: []int{
			2, 1, 5, 5, 9, 9, 0, 3, 2, 1, 5, 5, 9, 9, 0, 3, 2, 1, 5, 5, 9, 9, 0, 3, 2, 1, 5, 5, 9, 9, 0, 3,
			2, 1, 5, 5, 9, 9, 0, 3, 2, 1, 5, 5, 9, 9, 0, 3, 2, 1, 5, 5, 9, 9, 0, 3, 2, 1, 5, 5, 9, 9, 0, 3,
			2, 1, 5, 5, 9, 9, 0, 3, 2, 1, 5, 5, 9, 9, 0, 3, 2, 1, 5, 5, 9, 9, 0, 3, 2, 1, 5, 5, 9, 9, 0, 3,
			2, 1, 5, 5, 9, 9, 0, 3, 2, 1, 5, 5, 9, 9, 0, 3, 2, 1, 5, 5, 9, 9, 0, 3, 2, 1, 5, 5, 9, 9, 0, 3,
			2, 1, 5, 5, 9, 9, 0, 3, 2, 1, 5, 5, 9, 9, 0, 3, 2, 1, 5, 5, 9, 9, 0, 3, 2, 1, 5, 5, 9, 9, 0, 3,
			2, 1, 5, 5, 9, 9, 0, 3, 2, 1, 5, 5, 9, 9, 0, 3, 2, 1, 5, 5, 9, 9, 0, 3, 2, 1, 5, 5, 9, 9, 0, 3,
		},
		k: 5,
	},
}

func TestPartition(t *testing.T) {
	for _, c := range cases {
		head := toLinkedList(c.list)

		partition(head, c.k)

		if err := isValid(head, c.k, c.list); err != nil {
			t.Errorf("list: %v invalid result: %v err: %v", c.list, toSlice(head), err)
		}
	}
}

func TestPartitionStable(t *testing.T) {
	for _, c := range cases {
		head := toLinkedList(c.list)

		partitionStable(&head, c.k)

		if err := isValid(head, c.k, c.list); err != nil {
			t.Errorf("list: %v invalid result: %v err: %v", c.list, toSlice(head), err)
		}
	}
}

func BenchmarkPartition(b *testing.B) {
	for _, c := range cases {
		name := fmt.Sprintf("%v", c.list)
		if len(name) > 24 {
			name = name[:20] + "...]"
		}
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				head := toLinkedList(c.list)
				partition(head, c.k)
			}
		})
	}
}

func BenchmarkPartitionStable(b *testing.B) {
	for _, c := range cases {
		name := fmt.Sprintf("%v", c.list)
		if len(name) > 24 {
			name = name[:20] + "...]"
		}
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				head := toLinkedList(c.list)
				partitionStable(&head, c.k)
			}
		})
	}
}

func isValid(head *node, val int, orig []int) error {
	flip := false
	var list []int
	for n := head; n != nil; n = n.next {
		list = append(list, n.val)
		if n.val < val {
			if flip {
				return errors.New("out of sequence")
			}
		} else {
			flip = true
		}
	}

	if !sameElements(list, orig) {
		return errors.New("elements differ")
	}
	return nil
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

func sameElements(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	a2 := make([]int, len(a))
	b2 := make([]int, len(b))
	copy(a2, a)
	copy(b2, b)
	sort.Ints(a2)
	sort.Ints(b2)

	for i := 0; i < len(a2); i++ {
		if a2[i] != b2[i] {
			return false
		}
	}

	return true
}
