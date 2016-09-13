package delmidnode

import (
	"testing"
)

var cases = []struct {
	list      []string
	delIndex  int
	expected  []string
	expectErr bool
}{
	{
		list:     []string{"a", "b", "c", "d", "e", "f"},
		delIndex: 2,
		expected: []string{"a", "b", "d", "e", "f"},
	},
	{
		list:     []string{"a", "b", "c", "d", "e", "f"},
		delIndex: 1,
		expected: []string{"a", "c", "d", "e", "f"},
	},
	{
		list:     []string{"a", "b", "c", "d", "e", "f"},
		delIndex: 4,
		expected: []string{"a", "b", "c", "d", "f"},
	},
	{
		list:     []string{"a", "b", "c", "d", "e", "f"},
		delIndex: 0,
		expected: []string{"b", "c", "d", "e", "f"},
	},
	{
		list:      []string{"a", "b", "c", "d", "e", "f"},
		delIndex:  5,
		expectErr: true,
	},
}

func TestDelNode(t *testing.T) {
	for _, c := range cases {
		head := toLinkedList(c.list)
		cutNode := head
		for i := 0; i < c.delIndex; i++ {
			cutNode = cutNode.next
		}
		cutNodeVal := cutNode.val

		err := delNode(cutNode)
		if c.expectErr {
			if err == nil {
				t.Errorf("orig: %v delNode: %s want: err got: nil", c.list, cutNodeVal)
			}
			continue
		}

		if err != nil {
			t.Errorf("orig: %v delNode: %s err: %v", c.list, cutNodeVal, err)
			continue
		}
		actual := toSlice(head)
		if !equal(c.expected, actual) {
			t.Errorf("orig: %v delNode: %s want: %v got: %v", c.list, cutNodeVal, c.expected, actual)
		}
	}
}

func equal(a, b []string) bool {
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

func toLinkedList(list []string) *node {
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

func toSlice(head *node) []string {
	if head == nil {
		return nil
	}
	var list []string
	for n := head; n != nil; n = n.next {
		list = append(list, n.val)
	}
	return list
}
