package intersection

import (
	"testing"
)

func TestIntersection(t *testing.T) {
	a := getNodes(10, 100)
	b := getNodes(10, 1)
	c := getNodes(20, 1000)
	getTail(a).next = c
	getTail(b).next = c
	i := getIntersection(a, b)
	if i != nil {
		t.Logf("intsection at: %d", i.val)
	} else {
		t.Logf("no intersection")
	}
}

func getNodes(startVal, length int) *node {
	if length <= 0 {
		return nil
	}
	head := &node{val: startVal}
	for n, i := head, 1; i < length; i++ {
		n.next = &node{val: startVal + i}
		n = n.next
	}
	return head
}

func getTail(n *node) *node {
	for {
		if n.next == nil {
			return n
		}
		n = n.next
	}
}
