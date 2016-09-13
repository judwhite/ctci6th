package kthtolast

type node struct {
	next *node
	val  int
}

func find(head *node, k int) *node {
	if head == nil {
		return nil
	}
	kthToLast := head
	for n := head.next; n != nil; n = n.next {
		if k == 0 {
			kthToLast = kthToLast.next
		} else {
			k--
		}
	}
	if k == 0 {
		return kthToLast
	}
	return nil
}

// findBook finds the kth to last solution using the book solution
func findBook(head *node, k int) *node {
	p1, p2 := head, head

	// Note, the definition of k=1 is the last element isn't given in the original problem, it's only mentioned
	// in the solution. I changed the implemention here to "i <= k" from "i < k" to fit with k=0 being the last
	// element.

	// Move p1 k+1 nodes into the list.
	for i := 0; i <= k; i++ {
		if p1 == nil {
			return nil // Out of bounds
		}
		p1 = p1.next
	}

	for p1 != nil {
		p1 = p1.next
		p2 = p2.next
	}
	return p2
}
