package sumlists

func sumReverseBook(l1, l2 *node) *node {
	return addListsReverseCarry(l1, l2, 0)
}

func addListsReverseCarry(l1, l2 *node, carry int) *node {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	result := node{}
	value := carry
	if l1 != nil {
		value += l1.val
	}
	if l2 != nil {
		value += l2.val
	}

	result.val = value % 10 // second digit of number

	// recurse
	if l1 != nil || l2 != nil {
		var next1, next2 *node
		if l1 != nil {
			next1 = l1.next
		}
		if l2 != nil {
			next2 = l2.next
		}
		if value >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		more := addListsReverseCarry(next1, next2, carry)
		result.next = more
	}
	return &result
}

type partialSum struct {
	sum   *node
	carry int
}

func sumBook(l1, l2 *node) *node {
	len1 := length(l1)
	len2 := length(l2)

	// pad the shorter list with zeros - see note (1)
	if len1 < len2 {
		l1 = padList(l1, len2-len1)
	} else {
		l2 = padList(l2, len1-len2)
	}

	// add lists
	sum := addListHelper(l1, l2)

	// if there was a carry value left over, insert this at the front of the list.
	// otherwise, just return the linked list.
	if sum.carry == 0 {
		return sum.sum
	}
	result := insertBefore(sum.sum, sum.carry)
	return result
}

func addListHelper(l1, l2 *node) partialSum {
	if l1 == nil && l2 == nil {
		return partialSum{}
	}
	// add smaller digits recursively
	sum := addListHelper(l1.next, l2.next)

	// add carry to current data
	val := sum.carry + l1.val + l2.val

	// insert sum of current digits
	fullResult := insertBefore(sum.sum, val%10)

	// return sum so far, and the carry value
	sum.sum = fullResult
	sum.carry = val / 10
	return sum
}

// padList pads the list with zeros
func padList(n *node, padding int) *node {
	head := n
	for i := 0; i < padding; i++ {
		head = insertBefore(head, 0)
	}
	return head
}

// insertBefore inserts a node in the front of a linked list
func insertBefore(list *node, data int) *node {
	return &node{val: data, next: list}
}

func length(n *node) int {
	count := 0
	for ; n != nil; n = n.next {
		count++
	}
	return count
}
