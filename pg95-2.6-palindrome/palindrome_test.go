package palindrome

import (
	"testing"
)

var cases = []struct {
	list     []int
	expected bool
}{
	{list: []int{}, expected: true},
	{list: []int{1}, expected: true},
	{list: []int{1, 1}, expected: true},
	{list: []int{1, 2}, expected: false},
	{list: []int{1, 2, 3}, expected: false},
	{list: []int{1, 1, 2, 2, 1, 1}, expected: true},
	{list: []int{1, 1, 2, 1, 1}, expected: true},
	{list: []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 3, 3, 3, 2, 2, 1}, expected: true},
}

func TestIsPalindrome(t *testing.T) {
	for _, c := range cases {
		list := toLinkedList(c.list)
		actual := isPalindrome(list)
		if actual != c.expected {
			t.Errorf("list: %v want: %v got: %v", c.list, c.expected, actual)
		}
	}
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
