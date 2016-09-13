package partition

type node struct {
	next *node
	val  int
}

func partition(head *node, val int) {
	for p1 := head; p1 != nil; p1 = p1.next {
		if p1.val >= val {
			p2 := p1.next
			for ; p2 != nil; p2 = p2.next {
				if p2.val < val {
					p1.val, p2.val = p2.val, p1.val
					break
				}
			}
			if p2 == nil {
				break
			}
		}
	}
}

func partitionStable(head **node, val int) {
	var tailStart *node
	var tail *node

	var newHead *node
	var newHeadStart *node

	for n := *head; n != nil; n = n.next {
		if n.val >= val {
			if tail == nil {
				tail = n
				tailStart = n
			} else {
				tail.next = n
				tail = n
			}
		} else {
			if newHead == nil {
				newHead = n
				newHeadStart = n
			} else {
				newHead.next = n
				newHead = n
			}
		}
	}

	if newHeadStart == nil {
		newHeadStart = tailStart
	}
	if newHead != nil {
		newHead.next = tailStart
	}
	if tail != nil {
		tail.next = nil
	}

	*head = newHeadStart
}
