package kthtolast

import (
	"fmt"
	"testing"
)

var cases = []struct {
	head *node
	k    int
}{
	{head: toLinkedList(0), k: 0},
	{head: toLinkedList(1), k: 0},
	{head: toLinkedList(2), k: 0},
	{head: toLinkedList(2), k: 1},
	{head: toLinkedList(2), k: 2},
	{head: toLinkedList(10), k: 5},
	{head: toLinkedList(100), k: 99},
	{head: toLinkedList(100), k: 0},
	{head: toLinkedList(10000), k: 0},
	{head: toLinkedList(10000), k: 5000},
	{head: toLinkedList(10000), k: 9999},
	{head: toLinkedList(1000000), k: 0},
	{head: toLinkedList(1000000), k: 500000},
	{head: toLinkedList(1000000), k: 999999},
}

func TestFind(t *testing.T) {
	for _, c := range cases {
		test(t, c.head, c.k, find)
	}
}

func TestFindBook(t *testing.T) {
	for _, c := range cases {
		test(t, c.head, c.k, findBook)
	}
}

func test(t *testing.T, head *node, k int, testFunc func(*node, int) *node) {
	list := toSlice(head)
	expectedIndex := len(list) - k - 1
	actual := testFunc(head, k)
	if expectedIndex < 0 {
		if actual != nil {
			t.Errorf("len(list): %d k: %d want: nil got: %d", len(list), k, actual.val)
		}
	} else {
		expected := list[expectedIndex]
		if actual == nil {
			t.Errorf("len(list): %d k: %d want: %d got: nil", len(list), k, expected.val)
		} else if actual != expected {
			t.Errorf("len(list): %d k: %d want: %d got: %d", len(list), k, expected.val, actual.val)
		}
	}
}

func BenchmarkFind(b *testing.B) {
	for _, c := range cases {
		list := toSlice(c.head)
		name := fmt.Sprintf("%d_%d", len(list), c.k)
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				find(c.head, c.k)
			}
		})
	}
}

func BenchmarkFindBook(b *testing.B) {
	for _, c := range cases {
		list := toSlice(c.head)
		name := fmt.Sprintf("%d_%d", len(list), c.k)
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				findBook(c.head, c.k)
			}
		})
	}
}

func toLinkedList(k int) *node {
	if k == 0 {
		return nil
	}
	head := &node{val: 0}
	for n, i := head, 1; i < k; i++ {
		n.next = &node{val: i}
		n = n.next
	}
	return head
}

func toSlice(head *node) []*node {
	var list []*node
	for n := head; n != nil; n = n.next {
		list = append(list, n)
	}
	return list
}
