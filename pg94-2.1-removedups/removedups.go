package removedups

type node struct {
	next *node
	val  int
}

func removeDupes(head *node) {
	if head == nil || head.next == nil {
		return
	}
	m := make(map[int]struct{})
	m[head.val] = struct{}{}

	prev := head
	for n := head.next; n != nil; n = n.next {
		_, exists := m[n.val]
		if exists {
			prev.next = n.next
		} else {
			m[n.val] = struct{}{}
			prev = n
		}
	}
}

func removeDupesNoBuffer(head *node) {
	if head == nil || head.next == nil {
		return
	}
	for n := head; n != nil; n = n.next {
		prev := n
		for n2 := n.next; n2 != nil; n2 = n2.next {
			if n.val == n2.val {
				prev.next = n2.next
			} else {
				prev = n2
			}
		}
	}
}
